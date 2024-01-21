package config

type Config struct {
	Host        string `env:"SERVER_HOST" env-default:"localhost"`
	Port        string `env:"SERVER_PORT" env-default:"9999"`
	DatabaseURL string `env:"DB_URL"`
}

func New(host, port, databaseURL, logLevel string) *Config {
	return &Config{
		Host:        host,
		Port:        port,
		DatabaseURL: databaseURL,
	}
}
