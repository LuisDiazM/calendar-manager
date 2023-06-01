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
	repository_zoom := new(mocks.ZoomApiRepository)
	var meeting entities.Meetings
	var token *entities.AccessTokenResponse = &entities.AccessTokenResponse{AccessToken: "123"}

	var meetingZoom *entities.MeetingResponse = &entities.MeetingResponse{Type: 2, StartTime: "0001-01-01T00:00:00", Timezone: "America/Bogota"}
	var meetingData entities.MeetingResponse = entities.MeetingResponse{Type: 2, StartTime: "0001-01-01T00:00:00", Timezone: "America/Bogota"}
	repository_db.On("CreateMeeting", mock.Anything).Return(nil)
	repository_zoom.On("GenerateAccessToken").Return(token)
	repository_zoom.On("CreateZoomMeeting", mock.Anything, mock.Anything, meetingData).Return(meetingZoom)

	usecase := NewCreateMeeting(repository_db, repository_zoom)
	notification := usecase.CreateMeeting(meeting)
	assert.NotNil(t, notification)
	repository_db.MethodCalled("CreateMeeting", mock.Anything)
	repository_zoom.MethodCalled("GenerateAccessToken")
	repository_zoom.MethodCalled("CreateZoomMeeting", token.AccessToken, mock.Anything, meetingData)

}

func TestUpdateMeeting(t *testing.T) {
	repository_db := new(mocks.WriteMeetingRepository)
	repository_zoom := new(mocks.ZoomApiRepository)
	var meeting entities.Meetings
	repository_db.On("UpdateMeeting", meeting).Return(nil)
	usecase := NewCreateMeeting(repository_db, repository_zoom)
	err := usecase.UpdateMeeting(meeting)
	assert.Nil(t, err)
	repository_db.MethodCalled("UpdateMeeting", meeting)
}

func TestDeleteMeeting(t *testing.T) {
	repository_db := new(mocks.WriteMeetingRepository)
	repository_zoom := new(mocks.ZoomApiRepository)
	repository_db.On("DeleteMeeting", mock.Anything).Return(nil)
	usecase := NewCreateMeeting(repository_db, repository_zoom)
	err := usecase.DeleteMeeting(mock.Anything)
	assert.Nil(t, err)
	repository_db.MethodCalled("DeleteMeeting", mock.Anything)
}
