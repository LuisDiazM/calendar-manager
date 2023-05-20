package usecases

import (
	"time"

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

func (usecase *ReadMeeting) FindMeetingsByUserId(userId string, timestamp time.Time) (*[]entities.Meetings, error) {
	return usecase.ReadMeetingRepo.GetMeetingsByUser(userId, timestamp)
}
