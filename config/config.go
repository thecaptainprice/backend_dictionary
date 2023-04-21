package config

type Config struct {
	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPass string
	Port   string
}

func Load() *Config {
	cfg := &Config{
		DBHost: "localhost",
		DBPort: "3306",
		DBName: "dictionary",
		DBUser: "root",
		DBPass: "",
		Port:   "3333",
	}
	return cfg
}
