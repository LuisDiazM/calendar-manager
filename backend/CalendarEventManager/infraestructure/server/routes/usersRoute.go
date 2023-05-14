package routes

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/app"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/server/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUsersRouter(group *gin.RouterGroup, app *app.Application) {
	group.POST("/users", controllers.CreateUser(app.UsersUsecase))
}
