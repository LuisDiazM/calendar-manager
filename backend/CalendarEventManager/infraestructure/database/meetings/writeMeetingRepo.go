package meetings

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/entities"
	meetingRepo "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/repositories"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/database"
)

type WriteMeetingRepo struct {
	database database.DatabaseGateway
}

func NewWriteMeetingRepository(database database.DatabaseGateway) meetingRepo.WriteMeetingRepository {
	return &WriteMeetingRepo{database: database}
}

func (repo *WriteMeetingRepo) CreateMeeting(meeting entities.Meetings) error {
	result := repo.database.DB().Create(&meeting)
	return result.Error
}

func (repo *WriteMeetingRepo) UpdateMeeting(meeting entities.Meetings) error {
	result := repo.database.DB().Model(&entities.Meetings{}).Where("id = ?", meeting.ID).Updates(meeting)
	return result.Error
}

func (repo *WriteMeetingRepo) DeleteMeeting(id string) error {
	result := repo.database.DB().Where("id = ?", id).Delete(&entities.Meetings{})
	return result.Error
}
