// Copyright 2019 The klaytn Authors
// This file is part of the klaytn library.
//
// The klaytn library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The klaytn library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the klaytn library. If not, see <http://www.gnu.org/licenses/>.

package reward

import (
	"github.com/klaytn/klaytn/blockchain"
	"github.com/klaytn/klaytn/blockchain/state"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/event"
	"github.com/klaytn/klaytn/params"
)

const (
	chainHeadChanSize = 100
)

// blockChain is an interface for blockchain.Blockchain used in reward package.
type blockChain interface {
	SubscribeChainHeadEvent(ch chan<- blockchain.ChainHeadEvent) event.Subscription
	GetBlockByNumber(number uint64) *types.Block
	StateAt(root common.Hash) (*state.StateDB, error)
	Config() *params.ChainConfig

	blockchain.ChainContext
}

type StakingManager struct {
	addressBookConnector *addressBookConnector
	stakingInfoCache     *stakingInfoCache
	stakingInfoDB        stakingInfoDB
	governanceHelper     governanceHelper
	blockchain           blockChain
	chainHeadChan        chan blockchain.ChainHeadEvent
	chainHeadSub         event.Subscription
}

func NewStakingManager(bc blockChain, gh governanceHelper, db stakingInfoDB) *StakingManager {
	return &StakingManager{
		addressBookConnector: newAddressBookConnector(bc, gh),
		stakingInfoCache:     newStakingInfoCache(),
		stakingInfoDB:        db,
		governanceHelper:     gh,
		blockchain:           bc,
		chainHeadChan:        make(chan blockchain.ChainHeadEvent, chainHeadChanSize),
	}
}

// GetStakingInfo returns a corresponding stakingInfo for a blockNum.
func (sm *StakingManager) GetStakingInfo(blockNum uint64) *StakingInfo {
	stakingBlockNumber := params.CalcStakingBlockNumber(blockNum)

	// Get staking info from cache
	if cachedStakingInfo := sm.stakingInfoCache.get(stakingBlockNumber); cachedStakingInfo != nil {
		logger.Debug("StakingInfoCache hit.", "blockNum", blockNum, "staking block number", stakingBlockNumber, "stakingInfo", cachedStakingInfo)
		return cachedStakingInfo
	}

	// Get staking info from DB
	if storedStakingInfo, err := sm.getStakingInfoFromDB(stakingBlockNumber); storedStakingInfo != nil && err == nil {
		logger.Debug("StakingInfoDB hit.", "blockNum", blockNum, "staking block number", stakingBlockNumber, "stakingInfo", storedStakingInfo)
		sm.stakingInfoCache.add(storedStakingInfo)
		return storedStakingInfo
	} else {
		logger.Warn("Failed to get stakingInfo from DB", "err", err, "blockNum", blockNum)
	}

	// Calculate staking info from block header and updates it to cache and db
	calcStakingInfo, _ := sm.updateStakingInfo(stakingBlockNumber)
	if calcStakingInfo == nil {
		logger.Error("Failed to update stakingInfo", "blockNum", blockNum, "staking block number", stakingBlockNumber)
		return nil
	}

	logger.Debug("Get stakingInfo from header.", "blockNum", blockNum, "staking block number", stakingBlockNumber, "stakingInfo", calcStakingInfo)
	return calcStakingInfo
}

// updateStakingInfo updates staking info in cache and db created from given block number.
func (sm *StakingManager) updateStakingInfo(blockNum uint64) (*StakingInfo, error) {
	stakingInfo, err := sm.addressBookConnector.getStakingInfoFromAddressBook(blockNum)
	if err != nil {
		return nil, err
	}

	sm.stakingInfoCache.add(stakingInfo)
	if err := sm.addStakingInfoToDB(stakingInfo); err != nil {
		logger.Warn("Failed to write staking info to db.", "err", err, "stakingInfo", stakingInfo)
		return stakingInfo, err
	}

	logger.Info("Add a new stakingInfo to stakingInfoCache and stakingInfoDB", "blockNum", blockNum)
	logger.Debug("Added stakingInfo", "stakingInfo", stakingInfo)
	return stakingInfo, nil
}

// Subscribe setups a channel to listen chain head event and starts a goroutine to update staking cache.
func (sm *StakingManager) Subscribe() {
	sm.chainHeadSub = sm.blockchain.SubscribeChainHeadEvent(sm.chainHeadChan)

	go sm.handleChainHeadEvent()
}

func (sm *StakingManager) handleChainHeadEvent() {
	defer sm.Unsubscribe()

	logger.Info("Start listening chain head event to update stakingInfoCache.")

	for {
		// A real event arrived, process interesting content
		select {
		// Handle ChainHeadEvent
		case ev := <-sm.chainHeadChan:
			if sm.governanceHelper.ProposerPolicy() == params.WeightedRandom {
				// check and update if staking info is not valid beforehand
				sm.GetStakingInfo(ev.Block.NumberU64())
			}
		case <-sm.chainHeadSub.Err():
			return
		}
	}
}

// Unsubscribe can unsubscribe a subscription to listen chain head event.
func (sm *StakingManager) Unsubscribe() {
	sm.chainHeadSub.Unsubscribe()
}
