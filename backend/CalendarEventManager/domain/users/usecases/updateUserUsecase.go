package usecases

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/users/entities"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/users/repositories"
)

type UpdateUser struct {
	UpdateUserRepo repositories.UpdateUserRepository
}

func NewUpdateUser(repo repositories.UpdateUserRepository) *UpdateUser {
	return &UpdateUser{
		UpdateUserRepo: repo}
}

func (usecase *UpdateUser) UpdateUser(user entities.Users) error {
	return usecase.UpdateUserRepo.UpdateUser(user)
}
