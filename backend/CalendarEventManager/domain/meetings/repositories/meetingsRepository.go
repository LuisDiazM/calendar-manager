package respositories

import (
	"time"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/entities"
)

type WriteMeetingRepository interface {
	CreateMeeting(meeting entities.Meetings) error
	UpdateMeeting(meeting entities.Meetings) error
	DeleteMeeting(id string) error
}

type ReadMeetingRepository interface {
	ReadMeeting(id string) (*entities.Meetings, error)
	GetMeetingsByUser(userId string, timestamp time.Time) (*[]entities.Meetings, error)
}

type ZoomApiRepository interface {
	GenerateAccessToken() *entities.AccessTokenResponse
	CreateZoomMeeting(token string, userId string, meeting entities.MeetingResponse) *entities.MeetingResponse
}
