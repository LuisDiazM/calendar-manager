package cmd

import (
	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/cmd/config"
	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/infraestructure/app"
	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/infraestructure/messaging"
	"github.com/google/wire"
)

var EnvironmentVariablesProvider = wire.NewSet(config.NewEnvironmentsSpecification)
var AppProvider = wire.NewSet(app.NewApplication)
var MessagingRabbitProvider = wire.NewSet(messaging.NewRabbitMQConsumer)
