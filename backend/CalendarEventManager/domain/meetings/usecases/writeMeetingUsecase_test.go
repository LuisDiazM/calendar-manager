package usecases

import (
	"testing"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/entities"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateMeeting(t *testing.T) {
	repository_db := new(mocks.WriteMeetingRepository)
	var meeting entities.Meetings

	repository_db.On("CreateMeeting", mock.Anything).Return(nil)

	usecase := NewCreateMeeting(repository_db)
	notification := usecase.CreateMeeting(meeting)
	assert.NotNil(t, notification)
	repository_db.MethodCalled("CreateMeeting", mock.Anything)

}

func TestUpdateMeeting(t *testing.T) {
	repository_db := new(mocks.WriteMeetingRepository)
	var meeting entities.Meetings
	repository_db.On("UpdateMeeting", meeting).Return(nil)
	usecase := NewCreateMeeting(repository_db)
	err := usecase.UpdateMeeting(meeting)
	assert.Nil(t, err)
	repository_db.MethodCalled("UpdateMeeting", meeting)
}

func TestDeleteMeeting(t *testing.T) {
	repository_db := new(mocks.WriteMeetingRepository)
	repository_db.On("DeleteMeeting", mock.Anything).Return(nil)
	usecase := NewCreateMeeting(repository_db)
	err := usecase.DeleteMeeting(mock.Anything)
	assert.Nil(t, err)
	repository_db.MethodCalled("DeleteMeeting", mock.Anything)
}
