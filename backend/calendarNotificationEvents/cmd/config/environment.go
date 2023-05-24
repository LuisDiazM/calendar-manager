package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	RABBITMQ_DEFAULT_USER  string `mapstructure:"RABBITMQ_DEFAULT_USER"`
	RABBITMQ_DEFAULT_PASS  string `mapstructure:"RABBITMQ_DEFAULT_PASS"`
	RABBITMQ_DEFAULT_VHOST string `mapstructure:"RABBITMQ_DEFAULT_VHOST"`
	RABBITMQ_HOST          string `mapstructure:"RABBITMQ_HOST"`
	APP_NAME               string `mapstructure:"APP_NAME"`
}

func NewEnvironmentsSpecification() *Env {
	var env Env
	viper.SetConfigFile("cmd/config/.env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	return &env
}
