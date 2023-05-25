package repositories

import (
	"context"

	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/domain/meetings/entities"
)

type MeetingsWriteRepository interface {
	CreateMeetingEvent(meeting entities.Meetings, ctx context.Context) interface{}
}
