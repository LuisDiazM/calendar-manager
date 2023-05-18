package routes

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/app"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/server/middlewares"
)

const (
	v1 = "/api/v1"
)

func SetUpRoutes(app *app.Application) {
	publicRoutesV1 := app.WebServer.Group(v1)
	RegisterHealthRouter(publicRoutesV1)

	privateRoutesV1 := app.WebServer.Group(v1)
	privateRoutesV1.Use(middlewares.JwtGoogle())
	RegisterUsersRouter(privateRoutesV1, app)
	RegisterMeetingsRouter(privateRoutesV1, app)

}
