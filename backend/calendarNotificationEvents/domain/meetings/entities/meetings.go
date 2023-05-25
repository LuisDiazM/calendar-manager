package entities

import "time"

type Meetings struct {
	ID                  string    `json:"id,omitempty" bson:"_id"`
	InvitedGuest        []string  `json:"invited_guest,omitempty" bson:"invitedGuest"`
	MeetingDate         time.Time `json:"meeting_date,omitempty" bson:"meetingDate"`
	Title               string    `json:"title,omitempty" bson:"title"`
	Description         string    `json:"description" bson:"description"`
	EventDuration       int8      `json:"event_duration,omitempty" bson:"eventDuration"`
	VideoConferenceLink string    `json:"video_conference_link,omitempty" bson:"videoConferenceLink"`
	UserID              string    `json:"user_id" bson:"userId"`
}
