package rpc

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func NewWebSocketCodec(conn *websocket.Conn, host string, req http.Header) ServerCodec {
	return newWebsocketCodec(conn, host, req)
}
