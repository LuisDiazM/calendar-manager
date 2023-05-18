package respositories

import "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/entities"

type CreateMeetingRepository interface {
	CreateMeeting(meeting entities.Meetings) error
}
