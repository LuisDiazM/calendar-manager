package usecases

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/usecases/users/repositories"
)

type UsersUsecase struct {
	CreateUserUsecase *CreateUser
}

func NewUsersUsecase(createUsersRepo repositories.CreateUserRepository) *UsersUsecase {
	return &UsersUsecase{CreateUserUsecase: NewCreateUser(createUsersRepo)}
}
