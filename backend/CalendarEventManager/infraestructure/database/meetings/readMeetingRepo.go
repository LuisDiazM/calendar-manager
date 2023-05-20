package meetings

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/entities"
	respositories "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/repositories"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/database"
)

type ReadMeetingRepo struct {
	database database.DatabaseGateway
}

func NewReadMeetingRepo(database database.DatabaseGateway) respositories.ReadMeetingRepository {
	return &ReadMeetingRepo{database: database}
}

func (repo *ReadMeetingRepo) ReadMeeting(id string) (*entities.Meetings, error) {
	var meeting entities.Meetings
	result := repo.database.DB().First(&meeting, "id = ?", id)
	err := result.Error
	if err != nil {
		return nil, err
	} else {
		return &meeting, nil
	}
}

func (repo *ReadMeetingRepo) GetMeetingsByUser(userId string) (*[]entities.Meetings, error) {
	var meetings []entities.Meetings
	result := repo.database.DB().Where(&entities.Meetings{UserID: userId}).Find(&meetings)
	if result.Error != nil {
		return nil, result.Error
	}
	return &meetings, nil
}
