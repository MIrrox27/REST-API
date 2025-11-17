package main

import (
	//"log" // логирование ошибок и информации
	//"net/http"
	"fmt"
	"path/filepath"

	"github.com/MIrrox27/REST-API/internal/adapter/http"
)

func main() {
	indexPath := filepath.Join("frontend", "templates", "index.html") // в будущем будем брать из конфигов
	fmt.Println(indexPath)
	http.Router(indexPath)
}
