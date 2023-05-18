package controllers

import (
	"net/http"
	"reflect"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/users/entities"
	users "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/users/usecases"
	"github.com/gin-gonic/gin"
)

func CreateUser(usersUsecase *users.UsersUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user entities.Users
		err := ctx.BindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
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

func UpdateUser(usersUsecase *users.UsersUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user entities.Users
		err := ctx.BindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		err = usersUsecase.UpdateUserUsecase.UpdateUser(user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}
}
