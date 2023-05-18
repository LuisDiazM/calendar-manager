package meetings

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/entities"
	meetingRepo "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/repositories"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/database"
)

type CreateMeeting struct {
	database database.DatabaseGateway
}

func NewCreateMeetingRepo(database database.DatabaseGateway) meetingRepo.CreateMeetingRepository {
	return &CreateMeeting{database: database}
}

func (repo *CreateMeeting) CreateMeeting(meeting entities.Meetings) error {
	result := repo.database.DB().Create(&meeting)
	return result.Error
}
