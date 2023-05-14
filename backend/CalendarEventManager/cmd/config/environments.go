package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	SERVER_PORT       int    `mapstructure:"SERVER_PORT"`
	POSTGRES_PASSWORD string `mapstructure:"POSTGRES_PASSWORD"`
	POSTGRES_USER     string `mapstructure:"POSTGRES_USER"`
	POSTGRES_DB       string `mapstructure:"POSTGRES_DB"`
	POSTGRES_PORT     string `mapstructure:"POSTGRES_PORT"`
	POSTGRES_HOST     string `mapstructure:"POSTGRES_HOST"`
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
