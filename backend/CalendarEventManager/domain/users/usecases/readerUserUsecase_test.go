package usecases

import (
	"testing"

	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/users/entities"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/users/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUsersUsecaseFindUserById(t *testing.T) {
	repository := new(mocks.GetUserRepository)
	var user entities.Users
	var userReponse = &user
	repository.On("GetUserById", mock.Anything).Return(userReponse, nil)
	usecase := NewGetUsersUsecase(repository)
	response, _ := usecase.FindUserById(mock.Anything)
	assert.Empty(t, response)
	repository.MethodCalled("GetUserById", mock.Anything)
}

func TestGetUsersUsecaseFindUserByINil(t *testing.T) {
	repository := new(mocks.GetUserRepository)
	repository.On("GetUserById", mock.Anything).Return(nil, nil)
	usecase := NewGetUsersUsecase(repository)
	response, _ := usecase.FindUserById(mock.Anything)
	assert.Empty(t, response)
	repository.MethodCalled("GetUserById", mock.Anything)
}
