package app

import (
	"fmt"

	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/cmd/config"
	meetingsUsecase "github.com/LuisDiazM/calendar-manager/calendar-notification-events/domain/meetings"
	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/infraestructure/database"
	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/infraestructure/messaging"
)

type Application struct {
	Env             *config.Env
	MessagingRabbit *messaging.RabbitMQConsumer
	Database        *database.DatabaseImp
	MeetingUsecase  *meetingsUsecase.MeetingsUsecase
}

func NewApplication(configVars *config.Env,
	messaginRabbit *messaging.RabbitMQConsumer,
	database *database.DatabaseImp,
	meetingUsecase *meetingsUsecase.MeetingsUsecase) *Application {
	return &Application{Env: configVars,
		MessagingRabbit: messaginRabbit,
		Database:        database,
		MeetingUsecase:  meetingUsecase}
}

func (app Application) RegistryEventHandlers() {
	app.MessagingRabbit.AddEventHandler("registryMeeting", app.CreateMeetingHandler)
}

func (app Application) Start() {
	fmt.Println("start events consumer...")
	app.RegistryEventHandlers()
	app.MessagingRabbit.Start()
}
