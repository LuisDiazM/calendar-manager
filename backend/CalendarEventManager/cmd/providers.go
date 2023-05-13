package cmd

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/cmd/config"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/app"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/server"
	"github.com/google/wire"
)

var EnvironmentVariablesProvider = wire.NewSet(config.NewEnvironmentsSpecification)
var AppProvider = wire.NewSet(app.NewApplication)
var HTTPServerProvider = wire.NewSet(server.NewHTTPServer)
