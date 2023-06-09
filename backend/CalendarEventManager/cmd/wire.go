//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/app"
	"github.com/google/wire"
)

func CreateApp() *app.Application {
	wire.Build(AppProvider,
		EnvironmentVariablesProvider,
		HTTPServerProvider,
		DatabaseProvider,
		BrokerProvider,
		ZoomApiProvider,

		UsersUsecasesProvider,
		MeetingUsecaseProvider,

		UsersRepositoryProvider,
		UsersRepositoryReaderProvider,
		UsersRepositoryUpdateProvider,
		MeetingRepositoryCreateProvider,
		MeetRepositoryReadProvider,
		ZoomMeetingRepositoryProvider,
	)
	return new(app.Application)
}
