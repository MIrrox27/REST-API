// файл с описанием структуры конфигурации приложения и функцией для её загрузки, нужен для централизованного управления настройками приложения
package config

import ( // импортируем необходимые пакеты
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string               `yaml:"env" env-default:"local"` // необходим для определения окружения приложения
	StoragePath string               `yaml:"storage_path"`            // путь до файла с хранилищем
	HTTPServer  `yaml:"http_server"` // адрес внешнего HTTP сервера
}

type HTTPServer struct {
	Address     string        `yaml:"address"`      // адрес для запуска HTTP сервера
	Timeout     time.Duration `yaml:"timeout"`      // таймаут для чтения и записи данных
	IdleTimeout time.Duration `yaml:"idle_timeout"` // таймаут для неактивных соединений

}

func MustLoad() Config { // функция для загрузки конфигурации приложения
	config := os.Getenv("CONFIG_PATH")
	if config == "" { // если переменная окружения не установлена - падаем
		log.Fatal("CONFIG_PATH env variable is not set")
	}
	if _, err := os.Stat(config); os.IsNotExist(err) { // если файл не существует - падаем
		log.Fatalf("config file does not exist: %s", config)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(config, &cfg); err != nil { // читаем конфиг из файла, если ошибка - падаем
		log.Fatalf("failed to read config: %v", err)
	}

	return cfg
}
