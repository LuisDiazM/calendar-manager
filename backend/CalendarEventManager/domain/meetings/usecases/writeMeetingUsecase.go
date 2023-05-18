package usecases

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/entities"
	respositories "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/repositories"
)

type WriteMeeting struct {
	createRepo respositories.WriteMeetingRepository
}

func NewCreateMeeting(createRepo respositories.WriteMeetingRepository) *WriteMeeting {
	return &WriteMeeting{
		createRepo: createRepo,
	}
}

func (usecase *WriteMeeting) CreateMeeting(meeting entities.Meetings) error {
	return usecase.createRepo.CreateMeeting(meeting)
}

func (usecase *WriteMeeting) UpdateMeeting(meeting entities.Meetings) error {
	return usecase.createRepo.UpdateMeeting(meeting)
}

func (usecase *WriteMeeting) DeleteMeeting(id string) error {
	return usecase.createRepo.DeleteMeeting(id)
}
