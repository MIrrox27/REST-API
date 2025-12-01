package http

// этот файл нужен для инициализации роутера
import (
	"log" // логирование ошибок и информации
	//"net/http" // требуется для типов http.Request и http.ResponseWriter

	//"path/filepath" // сборка безопасных путей файловой системы

	"github.com/gin-gonic/gin" // веб-фреймворк Gin
	//github.com/gorilla/websocket" // реализация WebSocket
)

func NewServer() { // в кпараметре получаем
	// Route	r setup code goes here

	r := gin.Default()

	Router(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
