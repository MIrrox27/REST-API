package postgress

// этот файл нуже для сохранения и чтения данных из бд

import (
	//"database/sql"
	"github.com/jmoiron/sqlx"
	//"github.com/lib/pq"
)

type PostgresChat interface {
	ReadMessages(message_id int, message string, user_id int)
	SaveMessages(message_id int, message string, user_id int)
}

type PostgresRepo struct {
	// DB - это наше СОЕДИНЕНИЕ с базой данных.
	// Это указатель (*sqlx.DB), который мы получаем при старте приложения.
	// Мы внедряем его (DI) сюда, чтобы эта структура могла выполнять SQL-запросы.
	DB *sqlx.DB
}

func ReadMessages() {

}

func SaveMessages() {

}
