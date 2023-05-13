package routes

import "github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/app"

func SetUpRoutes(app *app.Application) {
	publicRoutes := app.WebServer.Group("")
	RegisterHealthRouter(publicRoutes)
}
