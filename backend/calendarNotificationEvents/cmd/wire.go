//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/infraestructure/app"
	"github.com/google/wire"
)

func CreateApp() *app.Application {
	wire.Build(AppProvider,
		EnvironmentVariablesProvider,
		MessagingRabbitProvider,
		DatabaseProvider,
		MeetingsUsecaseProvider,
		MeetingWriteRepositoryProvider,
	)
	return new(app.Application)
}
