package usecases

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/usecases/users/entities"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/usecases/users/repositories"
)

type CreateUser struct {
	CreateUserRepo repositories.CreateUserRepository
}

func NewCreateUser(repo repositories.CreateUserRepository) *CreateUser {
	return &CreateUser{
		CreateUserRepo: repo,
	}
}

func (usecase *CreateUser) InsertUser(user entities.Users) error {
	return usecase.CreateUserRepo.CreateUser(user)
}
