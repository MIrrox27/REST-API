package http

// этот файл нужен для инициализации роутера
import (
	"log" // логирование ошибок и информации
	//"net/http" // требуется для типов http.Request и http.ResponseWriter

	//"path/filepath" // сборка безопасных путей файловой системы

	"github.com/gin-gonic/gin" // веб-фреймворк Gin
	//github.com/gorilla/websocket" // реализация WebSocket
)

func Router(indexPath string) {
	// Router setup code goes here

	r := gin.Default()

	httpRouter(indexPath, r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
