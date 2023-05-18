package usecases

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/entities"
	respositories "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/repositories"
)

type CreateMeeting struct {
	createRepo respositories.CreateMeetingRepository
}

func NewCreateMeeting(createRepo respositories.CreateMeetingRepository) *CreateMeeting {
	return &CreateMeeting{
		createRepo: createRepo,
	}
}

func (usecase *CreateMeeting) CreateMeeting(meeting entities.Meetings) error {
	return usecase.createRepo.CreateMeeting(meeting)
}
