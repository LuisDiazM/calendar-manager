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
	)
	return new(app.Application)
}
