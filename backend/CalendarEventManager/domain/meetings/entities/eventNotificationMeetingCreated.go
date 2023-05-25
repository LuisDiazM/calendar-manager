package entities

import "time"

type EventMeetingNotification struct {
	ID                  string    `json:"id,omitempty"`
	InvitedGuest        []string  `json:"invited_guest,omitempty"`
	MeetingDate         time.Time `json:"meeting_date,omitempty"`
	Title               string    `json:"title,omitempty"`
	Description         string    `json:"description,omitempty"`
	EventDuration       int8      `json:"event_duration,omitempty"`
	VideoConferenceLink string    `json:"video_conference_link,omitempty"`
	UserID              string    `json:"user_id,omitempty"`
	ZoomMeetingId       int64     `json:"zoom_meeting_id,omitempty"`
}
