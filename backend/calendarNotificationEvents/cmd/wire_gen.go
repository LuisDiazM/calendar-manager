// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/cmd/config"
	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/infraestructure/app"
	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/infraestructure/messaging"
)

// Injectors from wire.go:

func CreateApp() *app.Application {
	env := config.NewEnvironmentsSpecification()
	rabbitMQConsumer := messaging.NewRabbitMQConsumer(env)
	application := app.NewApplication(env, rabbitMQConsumer)
	return application
}
