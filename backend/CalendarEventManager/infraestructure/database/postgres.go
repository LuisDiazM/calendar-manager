package database

import (
	"fmt"
	"log"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/cmd/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseImp struct {
	Client *gorm.DB
}

func NewDatabaseImp() DatabaseGateway {
	return &DatabaseImp{}
}

func (database *DatabaseImp) SetUp(environment config.Env) {
	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Bogota`, environment.POSTGRES_HOST, environment.POSTGRES_USER, environment.POSTGRES_PASSWORD, environment.POSTGRES_DB, environment.POSTGRES_PORT)
	var err error
	database.Client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	} else {
		log.Default().Println("database ready")
	}
}

func (database *DatabaseImp) DB() *gorm.DB {
	return database.Client
}
