package meetings

import (
	"context"

	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/domain/meetings/entities"
	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/domain/meetings/repositories"
	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/infraestructure/database"
)

type MeetingsWriteRepository struct {
	Db *database.DatabaseImp
}

func NewMeetingWriteRepository(db *database.DatabaseImp) repositories.MeetingsWriteRepository {
	return &MeetingsWriteRepository{Db: db}
}

func (repository *MeetingsWriteRepository) CreateMeetingEvent(meeting entities.Meetings, ctx context.Context) interface{} {
	col := repository.Db.Collection(database.DatabaseName, meetingsCollection)
	result := repository.Db.InsertOne(col, &ctx, meeting)
	return result.InsertedID
}
