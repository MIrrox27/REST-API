package websocket

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Hub *Hub
	// conn - это фактическое соединение WebSockets
	conn *websocket.Conn

	// send - это канал, через который Hub будет отправлять сообщения обратно клиенту.
	// Если Hub хочет отправить клиенту сообщение, он пишет в этот канал.
	send chan []byte
}

// NewClient - конструктор для создания нового клиента
func NewClient(hub *Hub, conn *websocket.Conn) *Client {
	return &Client{
		Hub:  hub,
		conn: conn,
		// Создаем канал для этого клиента
		send: make(chan []byte, 256),
	}
}
