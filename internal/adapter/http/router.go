package http

// этот файл нужен для инициализации роутера и настройки маршрутов
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

func Router() {
	// Router setup code goes here

	r := gin.Default()

	//indexPath := filepath.Join("../../..", "frontend", "templates", "index.html")
	httpRouter(r)

	// Запускаем HTTP-сервер на порту 8080; при ошибке — логируем и завершаем программу.
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
