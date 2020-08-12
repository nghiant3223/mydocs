package auth

type Config struct {
	password string
}

func NewConfig(password string) Config {
	return Config{password: password}
}
