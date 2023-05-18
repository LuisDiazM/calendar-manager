package controllers

import (
	"net/http"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/entities"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/usecases"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateMeetingController(usecases *usecases.MeetingsUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var meeting entities.Meetings
		err := ctx.BindJSON(&meeting)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		id := uuid.New()
		meeting.ID = id.String()
		err = usecases.CreateMeeting.CreateMeeting(meeting)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}
		ctx.JSON(http.StatusCreated, nil)
	}
}

func UpdateMeetingController(usecases *usecases.MeetingsUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var meet entities.Meetings
		id := ctx.Param("id")
		err := ctx.BindJSON(&meet)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		meet.ID = id
		err = usecases.CreateMeeting.UpdateMeeting(meet)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, meet)
	}
}

func DeleteMeetingController(usecases *usecases.MeetingsUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		err := usecases.CreateMeeting.DeleteMeeting(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
		}
		ctx.JSON(http.StatusNoContent, nil)
	}
}
