package main1

import (
	"REST-API/internal/config"
	"fmt"
	"os"

	"log/slog"

	"github.com/gin-gonic/gin"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
	//envTest = "test"

)

func main() {
	os.Setenv("CONFIG_PATH", "C:\\Users\\Пользователь\\.Portfolio\\REST-API\\config\\local.yaml") // устанавливаем путь к моим конфигам (сделан для того чтобы я мог работать)
	fmt.Println("CONFIG_PATH:", os.Getenv("CONFIG_PATH"))

	cfg := config.MustLoad() // загружаем конфигурацию приложения
	//fmt.Printf("Config loaded: %+v\n", cfg) // выводим загруженную конфигурацию для проверки
	log := setupLogger(cfg.Env) // настраиваем логгер в зависимости от окружения из конфига

	r := gin.Default()

	log.Info("starting url shortener", slog.String("env", cfg.Env)) // логируем старт приложения с указанием окружения из конфига, нужно для понимания в каком окружении запущено приложение
	log.Debug("debug messages are enabled")

	// TODO: init logger

	//TODO: init storage

	//TODO: init router

	//TODO: start server
}

func setupLogger(env string) *slog.Logger { // функция для настройки логгера в зависимости от окружения
	var log *slog.Logger // создаем новый логгер

	switch env { // если окружение локальное - создаем логгер с текстовым хендлером, если дев - с JSON хендлером и уровнем debug, если прод - с JSON хендлером и уровнем info
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})) // создаем логгер с текстовым хендлером для локальной разработки, хендлер это то, что обрабатывает логи и выводит их в нужном формате
		// первый параметр это куда выводить логи, второй это опции хендлера - уровень логирования. символ & означает что мы передаем указатель на структуру HandlerOptions. на данном уровне логгер будет выводить все логи от debug и выше (debug, info, warn, error)

	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})) // создаем логгер с JSON хендлером для дев окружения, первый параметр это куда выводить логи, второй это опции хендлера - уровень логирования. символ & означает что мы передаем указатель на структуру HandlerOptions. на данном уровне логгер будет выводить все логи от debug и выше (debug, info, warn, error)

	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})) // создаем логгер с JSON хендлером для прод окружения, первый параметр это куда выводить логи, второй это опции хендлера - уровень логирования. символ & означает что мы передаем указатель на структуру HandlerOptions. на данном уровне логгер будет выводить все логи от info и выше (info, warn, error)
	}
	return log // возвращаем созданный логгер
}
