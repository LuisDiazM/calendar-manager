package respositories

import "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/entities"

type WriteMeetingRepository interface {
	CreateMeeting(meeting entities.Meetings) error
	UpdateMeeting(meeting entities.Meetings) error
	DeleteMeeting(id string) error
}

type ReadMeetingRepository interface {
	ReadMeeting(id string) (*entities.Meetings, error)
	GetMeetingsByUser(userId string) (*[]entities.Meetings, error)
}
