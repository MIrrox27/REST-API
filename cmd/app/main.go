package main

import (
	//"log" // логирование ошибок и информации
	//"net/http"
	"github.com/MIrrox27/REST-API/internal/adapter/http"
)

func main() {

	http.Router() // вылезает ошибка, потому что в main.go нет доступа к http.Router()

}
