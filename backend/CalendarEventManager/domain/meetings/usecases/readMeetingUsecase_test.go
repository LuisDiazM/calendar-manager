package usecases

import (
	"testing"
	"time"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/entities"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindMeetingById(t *testing.T) {
	repository := new(mocks.ReadMeetingRepository)
	repository.On("ReadMeeting", mock.Anything).Return(nil, nil)
	usecase := NewReadMeeting(repository)
	meeting, err := usecase.FindMeetingById(mock.Anything)
	assert.Nil(t, err)
	assert.Nil(t, meeting)
	repository.MethodCalled("ReadMeeting", mock.Anything)
}

func TestFindMeetingsByUserId(t *testing.T) {
	repository := new(mocks.ReadMeetingRepository)
	var meetings *[]entities.Meetings = &[]entities.Meetings{}
	repository.On("GetMeetingsByUser", mock.Anything, mock.Anything).Return(meetings, nil)
	usecase := NewReadMeeting(repository)
	_, err := usecase.FindMeetingsByUserId(mock.Anything, time.Now())
	assert.Nil(t, err)
	repository.MethodCalled("GetMeetingsByUser", mock.Anything, mock.Anything)
}
