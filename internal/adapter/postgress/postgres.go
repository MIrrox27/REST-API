package postgress

// этот файл нуже для сохранения и чтения данных из бд

import (
	//"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	//"github.com/lib/pq"
	//"github.com/MIrrox27/REST-API/internal/domain/chat"
	"github.com/gorilla/websocket"
)

type PostgresChat interface {
	ReadMessages(c *gin.Context, conn *websocket.Conn, message_id int, message string, user_id int) *error
	HendleMessage(c *gin.Context, conn *websocket.Conn)
	SaveMessages(message_id int, message string, user_id int)
}

type PostgresRepo struct {
	// DB - это наше СОЕДИНЕНИЕ с базой данных.
	// Это указатель (*sqlx.DB), который мы получаем при старте приложения.
	// Мы внедряем его (DI) сюда, чтобы эта структура могла выполнять SQL-запросы.
	DB *sqlx.DB
}

func SaveMessages() {

}

func HendleMessage() {

}

func ReadMessages(c *gin.Context, conn *websocket.Conn) error {
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			// Обычно клиент закрыл соединение.
			return err
		}
		// Отправляем то же сообщение обратно.
		if err = conn.WriteMessage(mt, message); err != nil {
			return err
		}
	}

}
