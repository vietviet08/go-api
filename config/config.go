package config

type Config struct {
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	Port string
	Host string
}

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func GetConfig() Config {
	return Config{
		Server: ServerConfig{
			Host: "",
			Port: "8080",
		},
		DB: DBConfig{
			Host:   "localhost",
			Port:   "3306",
			DBName: "demo_go",
		},
	}
}
