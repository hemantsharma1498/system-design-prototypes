package client

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	hub
	conn  *websocket.Conn
	conns []*websocket.Conn
}
