package http

import (

	//"net/http"
	"fmt"

	"github.com/MIrrox27/REST-API/internal/service"
	"github.com/gin-gonic/gin"
)

// этот файл нужен для настройки маршрутов

func Router(r *gin.Engine, h *service.ChatServiceImpl) *gin.Engine { // входящие параметры: r - роутер Gin (передадим через Router()), indexPath - путь к index.html
	// здесь будут настраиваться маршруты

	httpRoutMain(r)
	httpRoutWs(r)
	fmt.Println(h)
	return r
}
