// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/cmd/config"
	usecases2 "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/usecases"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/users/usecases"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/app"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/database"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/database/meetings"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/database/users"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/messaging"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/server"
)

// Injectors from wire.go:

func CreateApp() *app.Application {
	env := config.NewEnvironmentsSpecification()
	engine := server.NewHTTPServer()
	databaseGateway := database.NewDatabaseImp()
	createUserRepository := users.NewUsersRepository(databaseGateway)
	getUserRepository := users.NewUsersRepositoryReader(databaseGateway)
	updateUserRepository := users.NewUsersRepositoryUpdate(databaseGateway)
	usersUsecase := usecases.NewUsersUsecase(createUserRepository, getUserRepository, updateUserRepository)
	writeMeetingRepository := meetings.NewWriteMeetingRepository(databaseGateway)
	readMeetingRepository := meetings.NewReadMeetingRepo(databaseGateway)
	meetingsUsecase := usecases2.NewMeetingUsecases(writeMeetingRepository, readMeetingRepository)
	rabbitProducer := messaging.NewProducer(env)
	application := app.NewApplication(env, engine, databaseGateway, usersUsecase, meetingsUsecase, rabbitProducer)
	return application
}
