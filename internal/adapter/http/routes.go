package http

import (
	//"log"
	"net/http"

	"github.com/MIrrox27/REST-API/internal/adapter/websocket"
	"github.com/gin-gonic/gin"
)

func httpRoutMain(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) { // маршрут "/"
		c.JSON(http.StatusOK, gin.H{
			"message": "хахахаха эта страница не используется",
		})
	})
}

func httpRoutWs(r *gin.Engine) { // маршрут "/ws"
	r.GET("/ws", websocket.WebsocketUpgrader) // апгреидим до вебсокетов
}
