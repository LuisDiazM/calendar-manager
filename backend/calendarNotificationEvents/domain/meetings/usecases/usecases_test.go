package usecases

import (
	"context"
	"testing"

	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/domain/meetings/entities"
	"github.com/LuisDiazM/calendar-manager/calendar-notification-events/domain/meetings/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateNotification(t *testing.T) {
	repo := new(mocks.MeetingsWriteRepository)
	var meeting entities.Meetings
	usecase := NewMeetingsUsecase(repo)
	repo.On("CreateMeetingEvent", meeting, context.TODO()).Return(mock.Anything)
	err := usecase.CreateNotifications(meeting, context.TODO())
	assert.Nil(t, err)
	repo.MethodCalled("CreateMeetingEvent", meeting, context.TODO())

}
