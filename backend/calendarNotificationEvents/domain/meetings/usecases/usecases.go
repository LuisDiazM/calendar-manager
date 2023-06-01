package usecases

import (
	"context"
	"log"

	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/domain/meetings/entities"
	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/domain/meetings/repositories"
)

type MeetingsUsecase struct {
	MeetingsWriteRepo repositories.MeetingsWriteRepository
}

func NewMeetingsUsecase(meetingWriteRepo repositories.MeetingsWriteRepository) *MeetingsUsecase {
	return &MeetingsUsecase{MeetingsWriteRepo: meetingWriteRepo}
}

func (usecase *MeetingsUsecase) CreateNotifications(meeting entities.Meetings, ctx context.Context) error {
	id := usecase.MeetingsWriteRepo.CreateMeetingEvent(meeting, ctx)
	log.Println(id)
	return nil
}
