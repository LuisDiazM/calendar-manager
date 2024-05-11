package routes

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/app"
)

const (
	v1 = "/api/v1"
)

func SetUpRoutes(app *app.Application) {
	publicRoutesV1 := app.WebServer.Group(v1)
	RegisterHealthRouter(publicRoutesV1)
	RegisterUsersRouter(publicRoutesV1, app)
	RegisterMeetingsRouter(publicRoutesV1, app)

}
