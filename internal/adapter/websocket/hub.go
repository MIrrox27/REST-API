package websocket

// этот файл нужен для реализации веб-сокетов

import (
	//"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket" // реализация WebSocket
)

var upgrader = websocket.Upgrader{ // создаем апгрейдер
	// Установим размер буферов чтения/записи
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// CheckOrigin - КРИТИЧЕСКИ ВАЖНО для CORS/безопасности.
	// Если true, разрешаем соединение с любого домена.
	// В продакшене нужно проверять происхождение запроса!
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebsocketRoute(c *gin.Context) {} // функцция которая абгрейдит соединение HTTP до WebSocket и обрабатывает сообщения, gorutine
