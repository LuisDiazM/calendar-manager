package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	SERVER_PORT            int    `mapstructure:"SERVER_PORT"`
	POSTGRES_PASSWORD      string `mapstructure:"POSTGRES_PASSWORD"`
	POSTGRES_USER          string `mapstructure:"POSTGRES_USER"`
	POSTGRES_DB            string `mapstructure:"POSTGRES_DB"`
	POSTGRES_PORT          string `mapstructure:"POSTGRES_PORT"`
	POSTGRES_HOST          string `mapstructure:"POSTGRES_HOST"`
	RABBITMQ_DEFAULT_USER  string `mapstructure:"RABBITMQ_DEFAULT_USER"`
	RABBITMQ_DEFAULT_PASS  string `mapstructure:"RABBITMQ_DEFAULT_PASS"`
	RABBITMQ_DEFAULT_VHOST string `mapstructure:"RABBITMQ_DEFAULT_VHOST"`
	RABBITMQ_HOST          string `mapstructure:"RABBITMQ_HOST"`
	APP_NAME_NOTIFICATIONS string `mapstructure:"APP_NAME_NOTIFICATIONS"`
	ZOOM_CLIENT_ID         string `mapstructure:"ZOOM_CLIENT_ID"`
	ZOOM_CLIENT_SECRET     string `mapstructure:"ZOOM_CLIENT_SECRET"`
	ZOOM_URL               string `mapstructure:"ZOOM_URL"`
	ZOOM_ACCOUNT_ID        string `mapstructure:"ZOOM_ACCOUNT_ID"`
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
