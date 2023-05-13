package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/server/controllers"
)

func RegisterHealthRouter(group *gin.RouterGroup) {
	group.GET("/health", controllers.HealthController)
}
