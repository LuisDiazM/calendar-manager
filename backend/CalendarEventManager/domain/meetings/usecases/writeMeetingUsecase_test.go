package usecases

import (
	"testing"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/entities"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateMeeting(t *testing.T) {
	repository := new(mocks.WriteMeetingRepository)
	var meeting entities.Meetings
	repository.On("CreateMeeting", mock.Anything).Return(nil)
	usecase := NewCreateMeeting(repository)
	err := usecase.CreateMeeting(meeting)
	assert.Nil(t, err)
	repository.MethodCalled("CreateMeeting", mock.Anything)
}
