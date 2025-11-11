package main

import (
	//"fmt"
	//"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.cob/gorilla/websocket"
)

func main() {
	//TODO: развернуть http сервер с использованием Gin
	//TODO: сделать форму для принятия сообщений
	//TODO: сделать отправку сообщений в окно браузера с использованием WebSocket

	server := gin.Default()
	gin.SetMode(gin.DebugMode)

	/*tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		// обработка ошибки
		log.Fatalf("failed to parse template: %v", err)
	}*/

	server.GET("/a", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
