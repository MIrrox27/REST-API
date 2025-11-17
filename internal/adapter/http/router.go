package http

// этот файл нужен для инициализации роутера
import (
	"log"      // логирование ошибок и информации
	"net/http" // требуется для типов http.Request и http.ResponseWriter

	//"path/filepath" // сборка безопасных путей файловой системы

	"github.com/gin-gonic/gin"     // веб-фреймворк Gin
	"github.com/gorilla/websocket" // реализация WebSocket
)

var upgrader = websocket.Upgrader{
	// CheckOrigin позволяет принимать соединения с любых источников (удобно в локалке).
	// В продакшне стоит ограничивать по происхождению.
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Router(indexPath string) {
	// Router setup code goes here

	r := gin.Default()

	httpRouter(indexPath, r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
