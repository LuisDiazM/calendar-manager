package app

import (
	"context"
	"encoding/json"

	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/domain/meetings/entities"
)

func (app *Application) CreateMeetingHandler(ctx context.Context, request map[string]interface{}) error {
	var structRequest entities.Meetings
	data, err := json.Marshal(request)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &structRequest)
	if err != nil {
		return err
	}
	err = app.MeetingUsecase.CreateNotifications(structRequest, ctx)
	return err
}
