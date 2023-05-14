package controllers

import (
	"net/http"
	"reflect"

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

func GetUserById(usersUsecase *users.UsersUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		user, _ := usersUsecase.GetUsersUsecase.FindUserById(id)
		isEmpty := reflect.DeepEqual(user, entities.Users{})
		if isEmpty {
			ctx.JSON(http.StatusNoContent, nil)
			return
		} else {
			ctx.JSON(http.StatusOK, user)
		}
	}
}
