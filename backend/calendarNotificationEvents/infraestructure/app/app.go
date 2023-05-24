package app

import (
	"fmt"

	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/cmd/config"
	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/infraestructure/messaging"
)

type Application struct {
	Env             *config.Env
	MessagingRabbit *messaging.RabbitMQConsumer
}

func NewApplication(configVars *config.Env, messaginRabbit *messaging.RabbitMQConsumer) *Application {
	return &Application{Env: configVars, MessagingRabbit: messaginRabbit}
}

func (app Application) RegistryEventHandlers() {
	app.MessagingRabbit.AddEventHandler("registryMeeting", app.CreateMeetingHandler)
}

func (app Application) Start() {
	fmt.Println("start events consumer...")
	app.RegistryEventHandlers()
	app.MessagingRabbit.Start()
}
