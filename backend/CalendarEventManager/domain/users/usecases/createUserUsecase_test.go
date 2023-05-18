package usecases

import (
	"testing"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/users/entities"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/users/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUserInsertUser(t *testing.T) {
	repository := new(mocks.CreateUserRepository)
	var user entities.Users
	repository.On("CreateUser", user).Return(nil)
	usecase := NewCreateUser(repository)
	error := usecase.InsertUser(user)
	assert.Nil(t, error)
	repository.MethodCalled("CreateUser", mock.Anything)
}
