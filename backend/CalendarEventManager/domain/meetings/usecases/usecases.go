package usecases

import (
	respositories "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/repositories"
)

type MeetingsUsecase struct {
	CreateMeeting *WriteMeeting
	ReadMeeting   *ReadMeeting
}

func NewMeetingUsecases(createRepo respositories.WriteMeetingRepository,
	readMeetingRepo respositories.ReadMeetingRepository,
	createMeetingZoomRepo respositories.ZoomApiRepository) *MeetingsUsecase {
	return &MeetingsUsecase{CreateMeeting: NewCreateMeeting(createRepo, createMeetingZoomRepo),
		ReadMeeting: NewReadMeeting(readMeetingRepo)}
}
