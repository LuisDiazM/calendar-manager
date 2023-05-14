package routes

import "github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/app"

func SetUpRoutes(app *app.Application) {
	publicRoutesV1 := app.WebServer.Group("/api/v1")
	RegisterHealthRouter(publicRoutesV1)
	RegisterUsersRouter(publicRoutesV1, app)
}
