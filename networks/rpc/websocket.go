// Modifications Copyright 2018 The klaytn Authors
// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.
//
// This file is derived from rpc/websocket.go (2018/06/04).
// Modified and improved for the klaytn development.

package rpc

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	gorillaws "github.com/gorilla/websocket"

	"github.com/klaytn/klaytn/common"
	"golang.org/x/net/websocket"
	"gopkg.in/fatih/set.v0"

	fastws "github.com/clevergo/websocket"
	"github.com/valyala/fasthttp"
)

const (
	wsReadBuffer  = 1024
	wsWriteBuffer = 1024
)

// websocketJSONCodec is a custom JSON codec with payload size enforcement and
// special number parsing.
var websocketJSONCodec = websocket.Codec{
	// Marshal is the stock JSON marshaller used by the websocket library too.
	Marshal: func(v interface{}) ([]byte, byte, error) {
		msg, err := json.Marshal(v)
		return msg, websocket.TextFrame, err
	},
	// Unmarshal is a specialized unmarshaller to properly convert numbers.
	Unmarshal: func(msg []byte, payloadType byte, v interface{}) error {
		dec := json.NewDecoder(bytes.NewReader(msg))
		dec.UseNumber()

		return dec.Decode(v)
	},
}

var wsBufferPool = new(sync.Pool)

// WebsocketHandler returns a handler that serves JSON-RPC to WebSocket connections.
//
// allowedOrigins should be a comma-separated list of allowed origin URLs.
// To allow connections with any origin, pass "*".
func (srv *Server) WebsocketHandler(allowedOrigins []string) http.Handler {
	return websocket.Server{
		Handshake: wsHandshakeValidator(allowedOrigins),
		Handler: func(conn *websocket.Conn) {
			if atomic.LoadInt32(&srv.wsConnCount) >= MaxWebsocketConnections {
				return
			}
			atomic.AddInt32(&srv.wsConnCount, 1)
			wsConnCounter.Inc(1)
			defer func() {
				atomic.AddInt32(&srv.wsConnCount, -1)
				wsConnCounter.Dec(1)
			}()
			// Create a custom encode/decode pair to enforce payload size and number encoding
			conn.MaxPayloadBytes = common.MaxRequestContentLength
			if WebsocketReadDeadline != 0 {
				conn.SetReadDeadline(time.Now().Add(time.Duration(WebsocketReadDeadline) * time.Second))
			}
			if WebsocketWriteDeadline != 0 {
				conn.SetWriteDeadline(time.Now().Add(time.Duration(WebsocketWriteDeadline) * time.Second))
			}

			encoder := func(v interface{}) error {
				return websocketJSONCodec.Send(conn, v)
			}
			decoder := func(v interface{}) error {
				return websocketJSONCodec.Receive(conn, v)
			}
			srv.ServeCodec(NewCodec(conn, encoder, decoder), OptionMethodInvocation|OptionSubscriptions)
		},
	}
}

var upgrader = fastws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// GSWebsocketHandler returns a handler that serves JSON-RPC to WebSocket connections.
//
// allowedOrigins should be a comma-separated list of allowed origin URLs.
// To allow connections with any origin, pass "*".
func (s *Server) GSWebsocketHandler(allowedOrigins []string) http.Handler {
	var upgrader = gorillaws.Upgrader{
		ReadBufferSize:  wsReadBuffer,
		WriteBufferSize: wsWriteBuffer,
		WriteBufferPool: wsBufferPool,
		CheckOrigin:     GSwsHandshakeValidator(allowedOrigins),
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logger.Debug("WebSocket upgrade failed", "err", err)
			return
		}
		if atomic.LoadInt32(&s.wsConnCount) >= MaxWebsocketConnections {
			// Gorilla websocket uses websocket.WriteControl method to close connection
			conn.WriteControl(gorillaws.CloseMessage, gorillaws.FormatCloseMessage(gorillaws.CloseGoingAway, "unexpected EOF"), time.Now().Add(time.Second))
			return
		}
		atomic.AddInt32(&s.wsConnCount, 1)
		wsConnCounter.Inc(1)
		defer func() {
			atomic.AddInt32(&s.wsConnCount, -1)
			wsConnCounter.Dec(1)
		}()

		if WebsocketReadDeadline != 0 {
			conn.SetReadDeadline(time.Now().Add(time.Duration(WebsocketReadDeadline) * time.Second))
		}
		if WebsocketWriteDeadline != 0 {
			conn.SetWriteDeadline(time.Now().Add(time.Duration(WebsocketWriteDeadline) * time.Second))
		}

		//Create a custom encode/decode pair to enforce payload size and number encoding
		encoder := func(v interface{}) error {
			msg, err := json.Marshal(v)
			if err != nil {
				return err
			}
			err = conn.WriteMessage(websocket.TextFrame, msg)

			if err != nil {
				return err
			}
			return err
		}
		decoder := func(v interface{}) error {
			_, data, err := conn.ReadMessage()
			if err != nil {
				return err
			}
			dec := json.NewDecoder(bytes.NewReader(data))
			dec.UseNumber()
			return dec.Decode(v)
		}

		codec := newWebsocketCodec(conn, encoder, decoder)

		s.ServeCodec(codec, 0)
	})
}

func (srv *Server) FastWebsocketHandler(ctx *fasthttp.RequestCtx) {
	// TODO-Klaytn handle websocket protocol
	protocol := ctx.Request.Header.Peek("Sec-WebSocket-Protocol")
	if protocol != nil {
		ctx.Response.Header.Set("Sec-WebSocket-Protocol", string(protocol))
	}

	err := upgrader.Upgrade(ctx, func(conn *fastws.Conn) {
		if atomic.LoadInt32(&srv.wsConnCount) >= MaxWebsocketConnections {
			return
		}
		atomic.AddInt32(&srv.wsConnCount, 1)
		wsConnCounter.Inc(1)
		defer func() {
			atomic.AddInt32(&srv.wsConnCount, -1)
			wsConnCounter.Dec(1)
		}()
		if WebsocketReadDeadline != 0 {
			conn.SetReadDeadline(time.Now().Add(time.Duration(WebsocketReadDeadline) * time.Second))
		}
		if WebsocketWriteDeadline != 0 {
			conn.SetWriteDeadline(time.Now().Add(time.Duration(WebsocketWriteDeadline) * time.Second))
		}
		// Create a custom encode/decode pair to enforce payload size and number encoding
		encoder := func(v interface{}) error {
			msg, err := json.Marshal(v)
			if err != nil {
				return err
			}
			err = conn.WriteMessage(websocket.TextFrame, msg)
			if err != nil {
				return err
			}
			// return fastws.WriteJSON(conn, v)
			return err
		}
		decoder := func(v interface{}) error {
			_, data, err := conn.ReadMessage()
			if err != nil {
				return err
			}
			dec := json.NewDecoder(bytes.NewReader(data))
			dec.UseNumber()
			return dec.Decode(v)
			// return fastws.ReadJSON(conn, v)
		}

		reader := bufio.NewReaderSize(bytes.NewReader(ctx.Request.Body()), common.MaxRequestContentLength)
		srv.ServeCodec(NewCodec(&httpReadWriteNopCloser{reader, ctx.Response.BodyWriter()}, encoder, decoder), OptionMethodInvocation|OptionSubscriptions)
	})
	if err != nil {
		logger.Error("FastWebsocketHandler fail to upgrade message", "err", err)
		return
	}
}

// NewWSServer creates a new websocket RPC server around an API provider.
//
// Deprecated: use Server.WebsocketHandler
func NewWSServer(allowedOrigins []string, srv *Server) *http.Server {
	return &http.Server{
		Handler: srv.WebsocketHandler(allowedOrigins),
	}
}

// NewGSWSServer creates a new websocket RPC server around an API provider.
//
// Deprecated: use Server.WebsocketHandler
func NewGSWSServer(allowedOrigins []string, srv *Server) *http.Server {
	return &http.Server{
		Handler: srv.GSWebsocketHandler(allowedOrigins),
	}
}

func NewFastWSServer(allowedOrigins []string, srv *Server) *fasthttp.Server {
	upgrader.CheckOrigin = wsFastHandshakeValidator(allowedOrigins)

	// TODO-Klaytn concurreny default (256 * 1024), goroutine limit (8192)
	return &fasthttp.Server{
		Concurrency:        ConcurrencyLimit,
		MaxRequestBodySize: common.MaxRequestContentLength,
		Handler:            srv.FastWebsocketHandler,
	}
}

func wsFastHandshakeValidator(allowedOrigins []string) func(ctx *fasthttp.RequestCtx) bool {
	origins := set.New()
	allowAllOrigins := false

	for _, origin := range allowedOrigins {
		if origin == "*" {
			allowAllOrigins = true
		}
		if origin != "" {
			origins.Add(strings.ToLower(origin))
		}
	}

	// allow localhost if no allowedOrigins are specified.
	if len(origins.List()) == 0 {
		origins.Add("http://localhost")
		if hostname, err := os.Hostname(); err == nil {
			origins.Add("http://" + strings.ToLower(hostname))
		}
	}

	logger.Debug(fmt.Sprintf("Allowed origin(s) for WS RPC interface %v\n", origins.List()))

	f := func(ctx *fasthttp.RequestCtx) bool {
		origin := strings.ToLower(string(ctx.Request.Header.Peek("Origin")))
		if allowAllOrigins || origins.Has(origin) {
			return true
		}
		logger.Warn(fmt.Sprintf("origin '%s' not allowed on WS-RPC interface\n", origin))
		return false
	}

	return f
}

// wsHandshakeValidator returns a handler that verifies the origin during the
// websocket upgrade process. When a '*' is specified as an allowed origins all
// connections are accepted.
func wsHandshakeValidator(allowedOrigins []string) func(*websocket.Config, *http.Request) error {
	origins := set.New()
	allowAllOrigins := false

	for _, origin := range allowedOrigins {
		if origin == "*" {
			allowAllOrigins = true
		}
		if origin != "" {
			origins.Add(strings.ToLower(origin))
		}
	}

	// allow localhost if no allowedOrigins are specified.
	if len(origins.List()) == 0 {
		origins.Add("http://localhost")
		if hostname, err := os.Hostname(); err == nil {
			origins.Add("http://" + strings.ToLower(hostname))
		}
	}

	logger.Debug(fmt.Sprintf("Allowed origin(s) for WS RPC interface %v\n", origins.List()))

	f := func(cfg *websocket.Config, req *http.Request) error {
		origin := strings.ToLower(req.Header.Get("Origin"))
		if allowAllOrigins || origins.Has(origin) {
			return nil
		}
		logger.Warn(fmt.Sprintf("origin '%s' not allowed on WS-RPC interface\n", origin))
		return fmt.Errorf("origin %s not allowed", origin)
	}

	return f
}

// GSwsHandshakeValidator returns a handler that verifies the origin during the
// websocket upgrade process. When a '*' is specified as an allowed origins all
// connections are accepted.
func GSwsHandshakeValidator(allowedOrigins []string) func(*http.Request) bool {
	origins := set.New()
	allowAllOrigins := false

	for _, origin := range allowedOrigins {
		if origin == "*" {
			allowAllOrigins = true
		}
		if origin != "" {
			origins.Add(origin)
		}
	}
	// allow localhost if no allowedOrigins are specified.
	if len(origins.List()) == 0 {
		origins.Add("http://localhost")
		if hostname, err := os.Hostname(); err == nil {
			origins.Add("http://" + hostname)
		}
	}
	logger.Debug(fmt.Sprintf("Allowed origin(s) for WS RPC interface %v", origins.List()))

	f := func(req *http.Request) bool {
		// Skip origin verification if no Origin header is present. The origin check
		// is supposed to protect against browser based attacks. Browsers always set
		// Origin. Non-browser software can put anything in origin and checking it doesn't
		// provide additional security.
		if _, ok := req.Header["Origin"]; !ok {
			return true
		}
		// Verify origin against allow list.
		origin := strings.ToLower(req.Header.Get("Origin"))
		if allowAllOrigins || origins.Has(origin) {
			return true
		}
		logger.Warn("Rejected WebSocket connection", "origin", origin)
		return false
	}

	return f
}

// DialWebsocket creates a new RPC client that communicates with a JSON-RPC server
// that is listening on the given endpoint.
//
// The context is used for the initial connection establishment. It does not
// affect subsequent interactions with the client.
func DialWebsocket(ctx context.Context, endpoint, origin string) (*Client, error) {
	if origin == "" {
		var err error
		if origin, err = os.Hostname(); err != nil {
			return nil, err
		}
		if strings.HasPrefix(endpoint, "wss") {
			origin = "https://" + strings.ToLower(origin)
		} else {
			origin = "http://" + strings.ToLower(origin)
		}
	}
	endpoint, header, err := wsClientHeaders(endpoint, origin)
	if err != nil {
		return nil, err
	}
	return NewGorillaWSClient(ctx, func(ctx context.Context) (*gorillaws.Conn, error) {
		dialer := gorillaws.Dialer{
			ReadBufferSize:  wsReadBuffer,
			WriteBufferSize: wsWriteBuffer,
			WriteBufferPool: wsBufferPool,
		}

		conn, resp, err := dialer.Dial(endpoint, header)
		//conn, _, err := dialer.Dial(endpoint, header)
		_ = resp
		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}
		return conn, err
	})
}

func wsClientHeaders(endpoint, origin string) (string, http.Header, error) {
	endpointURL, err := url.Parse(endpoint)
	if err != nil {
		return endpoint, nil, err
	}
	header := make(http.Header)
	if origin != "" {
		header.Add("origin", origin)
	}
	if endpointURL.User != nil {
		b64auth := base64.StdEncoding.EncodeToString([]byte(endpointURL.User.String()))
		header.Add("authorization", "Basic "+b64auth)
		endpointURL.User = nil
	}
	return endpointURL.String(), header, nil
}

func wsDialContext(ctx context.Context, config *websocket.Config) (*websocket.Conn, error) {
	var conn net.Conn
	var err error
	switch config.Location.Scheme {
	case "ws":
		conn, err = dialContext(ctx, "tcp", wsDialAddress(config.Location))
	case "wss":
		dialer := contextDialer(ctx)
		conn, err = tls.DialWithDialer(dialer, "tcp", wsDialAddress(config.Location), config.TlsConfig)
	default:
		err = websocket.ErrBadScheme
	}
	if err != nil {
		return nil, err
	}
	ws, err := websocket.NewClient(config, conn)
	if err != nil {
		conn.Close()
		return nil, err
	}
	return ws, err
}

var wsPortMap = map[string]string{"ws": "80", "wss": "443"}

func wsDialAddress(location *url.URL) string {
	if _, ok := wsPortMap[location.Scheme]; ok {
		if _, _, err := net.SplitHostPort(location.Host); err != nil {
			return net.JoinHostPort(location.Host, wsPortMap[location.Scheme])
		}
	}
	return location.Host
}

func dialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	d := &net.Dialer{KeepAlive: tcpKeepAliveInterval}
	return d.DialContext(ctx, network, addr)
}

func contextDialer(ctx context.Context) *net.Dialer {
	dialer := &net.Dialer{Cancel: ctx.Done(), KeepAlive: tcpKeepAliveInterval}
	if deadline, ok := ctx.Deadline(); ok {
		dialer.Deadline = deadline
	} else {
		dialer.Deadline = time.Now().Add(defaultDialTimeout)
	}
	return dialer
}

func newWebsocketCodec(conn *gorillaws.Conn, encode func(v interface{}) error, decode func(v interface{}) error) ServerCodec {
	conn.SetReadLimit(int64(common.MaxRequestContentLength))

	//return GSNewCodec(conn, conn.WriteJSON, conn.ReadJSON)
	return GSNewCodec(conn, encode, decode)
}
