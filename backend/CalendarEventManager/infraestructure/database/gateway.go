package database

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/cmd/config"
	"gorm.io/gorm"
)

type DatabaseGateway interface {
	SetUp(environment config.Env)
	DB() *gorm.DB
}
