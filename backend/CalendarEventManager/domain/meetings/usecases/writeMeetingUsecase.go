package usecases

import (
	"log"

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

func (usecase *WriteMeeting) CreateMeeting(meeting entities.Meetings) *entities.EventMeetingNotification {
	err := usecase.createRepo.CreateMeeting(meeting)
	if err != nil {
		log.Println(err)
		return nil
	}
	var responseEventMeeting entities.EventMeetingNotification = entities.EventMeetingNotification{
		ID:                  meeting.ID,
		InvitedGuest:        meeting.InvitedGuest,
		MeetingDate:         meeting.MeetingDate,
		Title:               meeting.Title,
		Description:         meeting.Description,
		EventDuration:       meeting.EventDuration,
		VideoConferenceLink: "",
		UserID:              meeting.UserID,
		ZoomMeetingId:       654654654,
	}
	return &responseEventMeeting
}

func (usecase *WriteMeeting) UpdateMeeting(meeting entities.Meetings) error {
	return usecase.createRepo.UpdateMeeting(meeting)
}

func (usecase *WriteMeeting) DeleteMeeting(id string) error {
	return usecase.createRepo.DeleteMeeting(id)
}
