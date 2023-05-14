package controllers

import (
	"net/http"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/usecases/users/entities"
	users "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/usecases/users/usecases"
	"github.com/gin-gonic/gin"
)

func CreateUser(usersUsecase *users.UsersUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user entities.Users
		err := ctx.BindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		err = usersUsecase.CreateUserUsecase.InsertUser(user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusCreated, user)
	}
}
