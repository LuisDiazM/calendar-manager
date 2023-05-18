package usecases

import respositories "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/repositories"

type MeetingsUsecase struct {
	CreateMeeting *WriteMeeting
}

func NewMeetingUsecases(createRepo respositories.WriteMeetingRepository) *MeetingsUsecase {
	return &MeetingsUsecase{CreateMeeting: NewCreateMeeting(createRepo)}
}
