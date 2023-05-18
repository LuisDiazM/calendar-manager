package usecases

import respositories "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/repositories"

type MeetingsUsecase struct {
	CreateMeeting *CreateMeeting
}

func NewMeetingUsecases(createRepo respositories.CreateMeetingRepository) *MeetingsUsecase {
	return &MeetingsUsecase{CreateMeeting: NewCreateMeeting(createRepo)}
}
