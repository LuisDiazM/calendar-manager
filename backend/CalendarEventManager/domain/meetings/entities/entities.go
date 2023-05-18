package entities

import (
	"time"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/users/entities"
	"gorm.io/datatypes"
)

type Meetings struct {
	ID                  string                      `json:"id,omitempty" gorm:"primaryKey"`
	InvitedGuest        datatypes.JSONSlice[string] `json:"invited_guest,omitempty"`
	MeetingDate         time.Time                   `json:"meeting_date,omitempty"`
	Title               string                      `json:"title,omitempty" gorm:"type:varchar(100)"`
	Description         string                      `json:"description" gorm:"type:varchar(300)"`
	EventDuration       int8                        `json:"event_duration,omitempty"`
	VideoConferenceLink string                      `json:"video_conference_link,omitempty" gorm:"type:varchar(100)"`
	UserID              string                      `json:"user_id"`
	Users               entities.Users              `gorm:"foreignKey:UserID"`
}
