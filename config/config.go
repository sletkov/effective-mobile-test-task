package config

type Config struct {
	Host        string
	Port        string
	DatabaseURL string
}

func New(host, port, databaseURL string) *Config {
	return &Config{
		Host:        host,
		Port:        port,
		DatabaseURL: databaseURL,
	}
}
