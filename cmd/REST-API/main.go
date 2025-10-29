package main

import (
	"REST-API/internal/config"
	"fmt"
)

func main() {

	cfg := config.MustLoad()

	fmt.Printf("Config loaded: %+v\n", cfg) // выводим загруженную конфигурацию для проверки
	// TODO: init config

	// TODO: init logger

	//TODO: init storage

	//TODO: init router

	//TODO: start server

}
