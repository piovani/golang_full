package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// API REST
	StageAPP    string `json:"STAGE_APP"`
	ApiRestPort int32  `envconfig:"API_REST_PORT"`

	// DB
	HostDB     string `envconfig:"DB_HOST"`
	PortDB     string `envconfig:"DB_PORT"`
	UserDB     string `envconfig:"DB_USER"`
	PasswordDB string `envconfig:"DB_PASSWORD"`
	DatabaseDB string `envconfig:"DB_DATABASE"`

	// AWS
	AwsAccessKeyID     string `envconfig:"AWS_ACCESS_KEY_ID"`
	AwsSecretAccessKey string `envconfig:"AWS_SECRET_ACCESS_KEY"`
	AwsRegion          string `envconfig:"AWS_REGION"`
	AwsBucket          string `envconfig:"AWS_BUCKET"`
	AwsPort            string `envconfig:"AWS_PORT"`
}

var Env Config = Config{}

func InitConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	return envconfig.Process("", &Env)
}
