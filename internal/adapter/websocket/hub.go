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

func WebsocketUpgrader(c *gin.Context) { // функцция которая абгрейдит соединение HTTP до WebSocket и обрабатывает сообщения, gorutine

	// Получаем *http.Request и http.ResponseWriter из контекста Gin.
	req := c.Request
	w := c.Writer

	// Выполняем апгрейд.
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		// Ошибка обычно происходит, если заголовки не соответствуют протоколу.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Обязательно закрываем соединение при выходе из функции.
	defer conn.Close()

	// Пример простого эхо‑сервера.
	for {
		// Читаем сообщение от клиента.
		mt, message, err := conn.ReadMessage()
		if err != nil {
			// Обычно клиент закрыл соединение.
			break
		}
		// Отправляем то же сообщение обратно.
		if err = conn.WriteMessage(mt, message); err != nil {
			break
		}
	}
}
