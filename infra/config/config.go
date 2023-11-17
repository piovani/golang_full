package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// API REST
	ApiRestPort int32 `envconfig:"API_REST_PORT"`

	// DB
	HostDB     string `envconfig:"DB_HOST"`
	PortDB     int    `envconfig:"DB_PORT"`
	UserDB     string `envconfig:"DB_USER"`
	PasswordDB string `envconfig:"DB_PASSWORD"`
	DatabaseDB string `envconfig:"DB_DATABASE"`
}

var Env Config = Config{}

func InitConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	return envconfig.Process("", &Env)
}
