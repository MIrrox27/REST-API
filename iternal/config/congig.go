 // файл с описанием структуры конфигурации приложения и функцией для её загрузки, нужен для централизованного управления настройками приложения
 pacage config

 type Config struct {
	Env         string       `yaml:"env"` // необходим для определения окружения приложения
	StoragePath string       `yaml:"storage_path"` // путь до файла с хранилищем
 }


 type HTTPServerConfig struct {
	Address     string `yaml:"address"`      // адрес для запуска HTTP сервера
	Timeout     string `yaml:"timeout"`      // таймаут для чтения и записи данных
	IdleTimeout string `yaml:"idle_timeout"` // таймаут для неактивных соединений
	HTTPServer string `yaml:"http_server"`  // адрес внешнего HTTP сервера
 }
 
 func MustLoad() Config{ // функция для загрузки конфигурации приложения
	config := os.Getenv("CONFIG_PATH")
	if config == "" { // если переменная окружения не установлена - падаем
		logger.Fatal("CONFIG_PATH env variable is not set")
	}
	if -, err := os.Stat(config); os.IsNotExist(err) { // если файл не существует - падаем
		logger.Fatalf("config file does not exist: %s", config)
	}

	var cfg Config 
	if err := cleanenv.ReadConfig(config, &cfg); err != nil { // читаем конфиг из файла, если ошибка - падаем
		logger.Fatalf("failed to read config: %v", err)
	}
 }