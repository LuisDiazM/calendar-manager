package usecases

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/entities"
	repositories "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/repositories"
)

type ReadMeeting struct {
	ReadMeetingRepo repositories.ReadMeetingRepository
}

func NewReadMeeting(readMeetingRepo repositories.ReadMeetingRepository) *ReadMeeting {
	return &ReadMeeting{
		ReadMeetingRepo: readMeetingRepo,
	}
}

func (usecase *ReadMeeting) FindMeetingById(id string) (*entities.Meetings, error) {
	return usecase.ReadMeetingRepo.ReadMeeting(id)
}
