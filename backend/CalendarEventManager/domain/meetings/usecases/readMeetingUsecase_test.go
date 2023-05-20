package usecases

import (
	"testing"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/meetings/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindMeetingById(t *testing.T) {
	repository := new(mocks.ReadMeetingRepository)
	repository.On("ReadMeeting", mock.Anything).Return(nil, nil)
	usecase := NewReadMeeting(repository)
	meeting, _ := usecase.FindMeetingById(mock.Anything)
	assert.Nil(t, meeting)
	repository.MethodCalled("ReadMeeting", mock.Anything)
}

func TestFindMeetingsByUserId(t *testing.T) {
	repository := new(mocks.ReadMeetingRepository)
	repository.On("GetMeetingsByUser", mock.Anything).Return(nil, nil)
	usecase := NewReadMeeting(repository)
	meeting, _ := usecase.FindMeetingsByUserId(mock.Anything)
	assert.Nil(t, meeting)
	repository.MethodCalled("GetMeetingsByUser", mock.Anything)
}
