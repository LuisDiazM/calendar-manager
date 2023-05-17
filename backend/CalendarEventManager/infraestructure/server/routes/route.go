package routes

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/app"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/server/middlewares"
)

func SetUpRoutes(app *app.Application) {
	publicRoutesV1 := app.WebServer.Group("/api/v1")
	privateRoutesV1 := app.WebServer.Group("/api/v1")
	privateRoutesV1.Use(middlewares.JwtGoogle())

	RegisterHealthRouter(publicRoutesV1)
	RegisterUsersRouter(privateRoutesV1, app)
}
