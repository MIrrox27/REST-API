package main

import (
	"REST-API/internal/config"
	"fmt"
	"os"
)

func main() {
	os.Setenv("CONFIG_PATH", "C:\\Users\\Пользователь\\.Portfolio\\REST-API\\config\\local.yaml") // устанавливаем путь к моим конфигам (сделан для того чтобы я мог работать)
	cfg := config.MustLoad()

	fmt.Println("CONFIG_PATH:", os.Getenv("CONFIG_PATH"))

	fmt.Printf("Config loaded: %+v\n", cfg) // выводим загруженную конфигурацию для проверки

	// TODO: init config

	// TODO: init logger

	//TODO: init storage

	//TODO: init router

	//TODO: start server

}
