package routes

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/app"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/server/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterMeetingsRouter(group *gin.RouterGroup, app *app.Application) {
	group.POST("/meetings", controllers.CreateMeetingController(app.MeetingsUsecase))
	group.PUT("/meeting/:id", controllers.UpdateMeetingController(app.MeetingsUsecase))
	group.DELETE("/meeting/:id", controllers.DeleteMeetingController(app.MeetingsUsecase))
	group.GET("/meeting/:id", controllers.ReadByIdMeetingController(app.MeetingsUsecase))
	group.GET("/meetings/user/:id", controllers.FindMeetingsByUserController(app.MeetingsUsecase))
}
