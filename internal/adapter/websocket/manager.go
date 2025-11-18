package websocket

// этот файл нужен для реализации веб-сокетов

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket" // реализация WebSocket
)

var upgrader = websocket.Upgrader{
	// CheckOrigin позволяет принимать соединения с любых источников (удобно в локалке).
	// В продакшне стоит ограничивать по происхождению.
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WebsocketRoute(c *gin.Context) {

	w := c.Writer    //  это http.ResponseWriter, нужен чтобы апгрейдить соединение
	req := c.Request // это *http.Request, тоже нужен для апгрейда

	conn, err := upgrader.Upgrade(w, req, nil) // апгрейдим соединение до WebSocket, третий параметр — заголовки ответа (не нужны, поэтому nil)

	if err != nil { // если апгрейд не удался — логируем причину и завершаем обработчик
		log.Println("upgrade:", err)
		return
	}
	defer conn.Close() // Гарантированно закроем соединение при выходе из обработчика, чтобы не было утечек.

	for {
		// ReadMessage читает тип сообщения (mt) и его содержимое (message). mt - message type
		mt, message, err := conn.ReadMessage()
		if err != nil {
			// При ошибке чтения (например, клиент закрыл соединение) — логируем и выходим.
			log.Println("read:", err)
			break
		}
		// Логируем полученное сообщение на сервере.
		log.Printf("recv: %s", message)

		// Отправляем обратно то же сообщение с тем же типом (echo).
		if err = conn.WriteMessage(mt, message); err != nil {
			// При ошибке записи логируем и выходим из цикла.
			log.Println("write:", err)
			break
		}
	}

}
