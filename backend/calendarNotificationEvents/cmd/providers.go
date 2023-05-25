package cmd

import (
	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/cmd/config"
	meetingUsecase "github.com/LuisDiazM/calendar-manager/calendar-notification-events/domain/meetings"
	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/infraestructure/app"
	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/infraestructure/database"
	meetingsRepo "github.com/LuisDiazM/calendar-manager/calendar-notification-events/infraestructure/database/meetings"
	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/infraestructure/messaging"
	"github.com/google/wire"
)

var EnvironmentVariablesProvider = wire.NewSet(config.NewEnvironmentsSpecification)
var AppProvider = wire.NewSet(app.NewApplication)
var MessagingRabbitProvider = wire.NewSet(messaging.NewRabbitMQConsumer)
var DatabaseProvider = wire.NewSet(database.NewDatabaseImplementation)

// repositories
var MeetingWriteRepositoryProvider = wire.NewSet(meetingsRepo.NewMeetingWriteRepository)

// usecases
var MeetingsUsecaseProvider = wire.NewSet(meetingUsecase.NewMeetingsUsecase)
