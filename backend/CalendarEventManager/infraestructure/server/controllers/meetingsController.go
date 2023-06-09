package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/entities"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/usecases"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/app"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/messaging"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateMeetingController(app *app.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var meeting entities.Meetings
		err := ctx.BindJSON(&meeting)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		id := uuid.New()
		meeting.ID = id.String()
		eventMeetingNotification := app.MeetingsUsecase.CreateMeeting.CreateMeeting(meeting)
		if eventMeetingNotification == nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		} else {
			var event messaging.Event = messaging.Event{Name: "registryMeeting", EventId: uuid.NewString(), Data: eventMeetingNotification}
			err = app.BrokerProducer.PublishEvent(event, ctx.Request.Context())
			if err != nil {
				log.Println(err)
			}
			ctx.JSON(http.StatusCreated, eventMeetingNotification)
		}
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

func ReadByIdMeetingController(usecases *usecases.MeetingsUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		meeting, _ := usecases.ReadMeeting.FindMeetingById(id)
		if meeting != nil {
			ctx.JSON(http.StatusOK, meeting)
		} else {
			ctx.JSON(http.StatusNoContent, nil)
		}
	}
}

func FindMeetingsByUserController(usecases *usecases.MeetingsUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		after := ctx.Query("after")
		timestamp, err := time.Parse("2006-01-02T15:04:05.000Z", after)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		meetings, _ := usecases.ReadMeeting.FindMeetingsByUserId(id, timestamp)
		if len(*meetings) > 0 {
			ctx.JSON(http.StatusOK, meetings)
		} else {
			ctx.JSON(http.StatusNoContent, nil)
		}
	}
}
