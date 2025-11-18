package http

import (
	"log"
	//"net/http"

	"github.com/gin-gonic/gin"
)

// этот файл нужен для настройки маршрутов

func httpRouter(indexPath string, r *gin.Engine) { // входящие параметры: r - роутер Gin (передадим через Router()), indexPath - путь к index.html
	// здесь будут настраиваться маршруты
	log.Println(indexPath)
	r.LoadHTMLFiles(indexPath)

	httpRoutMain(r)
	httpRoutWs(r)

}
