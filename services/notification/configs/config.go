package configs

type Config struct {
	Port string
}

func NewConfig() *Config {
	return &Config{
		Port: ":8082",
	}
}
