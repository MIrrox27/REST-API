package test

import (
	"log"

	"github.com/gorilla/websocket"
)

func TestWebsocket() {
	// Обратите внимание на схему ws://
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatalf("dial error: %v", err)
	}
	defer conn.Close()

	// Отправляем тестовое сообщение
	if err = conn.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
		log.Fatalf("write error: %v", err)
	}

	// Читаем ответ (эхо‑сервер вернёт то же самое)
	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Fatalf("read error: %v", err)
	}
	log.Printf("received: %s", msg)
}
