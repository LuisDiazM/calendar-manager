// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/cmd/config"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/app"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/server"
)

// Injectors from wire.go:

func CreateApp() *app.Application {
	env := config.NewEnvironmentsSpecification()
	engine := server.NewHTTPServer()
	application := app.NewApplication(env, engine)
	return application
}
