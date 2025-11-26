package http

import (

	//"net/http"

	"github.com/gin-gonic/gin"
)

// этот файл нужен для настройки маршрутов

func Router(r *gin.Engine) { // входящие параметры: r - роутер Gin (передадим через Router()), indexPath - путь к index.html
	// здесь будут настраиваться маршруты

	httpRoutMain(r)
	httpRoutWs(r)

}
