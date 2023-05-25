package usecases

import (
	"log"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/entities"
	respositories "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/repositories"
)

type WriteMeeting struct {
	createRepo            respositories.WriteMeetingRepository
	createMeetingZoomRepo respositories.ZoomApiRepository
}

func NewCreateMeeting(createRepo respositories.WriteMeetingRepository, createMeetingZoomRepo respositories.ZoomApiRepository) *WriteMeeting {
	return &WriteMeeting{
		createRepo:            createRepo,
		createMeetingZoomRepo: createMeetingZoomRepo,
	}
}

func (usecase *WriteMeeting) CreateMeeting(meeting entities.Meetings) *entities.EventMeetingNotification {
	token := usecase.createMeetingZoomRepo.GenerateAccessToken()
	var meetingZoom entities.MeetingResponse = entities.MeetingResponse{
		Topic:     meeting.Description,
		Type:      2,
		StartTime: meeting.MeetingDate.Format("2006-01-02T15:04:05"),
		Duration:  int64(meeting.EventDuration),
		Timezone:  "America/Bogota",
		Agenda:    meeting.Title,
	}
	meetingZoomResponse := usecase.createMeetingZoomRepo.CreateZoomMeeting(token.AccessToken, "me", meetingZoom)
	meeting.VideoConferenceLink = meetingZoomResponse.JoinURL
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
		VideoConferenceLink: meetingZoomResponse.JoinURL,
		UserID:              meeting.UserID,
		ZoomMeetingId:       meetingZoomResponse.ID,
	}
	return &responseEventMeeting
}

func (usecase *WriteMeeting) UpdateMeeting(meeting entities.Meetings) error {
	return usecase.createRepo.UpdateMeeting(meeting)
}

func (usecase *WriteMeeting) DeleteMeeting(id string) error {
	return usecase.createRepo.DeleteMeeting(id)
}
