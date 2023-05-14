package usecases

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/usecases/users/repositories"
)

type UsersUsecase struct {
	CreateUserUsecase *CreateUser
	GetUsersUsecase   *GetUsersUsecase
}

func NewUsersUsecase(createUsersRepo repositories.CreateUserRepository, getUsersRepo repositories.GetUserRepository) *UsersUsecase {
	return &UsersUsecase{CreateUserUsecase: NewCreateUser(createUsersRepo),
		GetUsersUsecase: NewGetUsersUsecase(getUsersRepo)}
}
