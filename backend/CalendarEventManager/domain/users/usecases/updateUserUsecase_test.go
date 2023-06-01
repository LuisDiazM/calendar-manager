package usecases

import (
	"testing"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/users/entities"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/users/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateUser(t *testing.T) {
	repository := new(mocks.UpdateUserRepository)
	var user entities.Users
	repository.On("UpdateUser", user).Return(nil)
	usecase := NewUpdateUser(repository)
	error := usecase.UpdateUser(user)
	assert.Nil(t, error)
	repository.MethodCalled("UpdateUser", mock.Anything)
}
