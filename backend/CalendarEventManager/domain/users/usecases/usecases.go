package usecases

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/users/repositories"
)

type UsersUsecase struct {
	CreateUserUsecase *CreateUser
	GetUsersUsecase   *GetUsersUsecase
	UpdateUserUsecase *UpdateUser
}

func NewUsersUsecase(createUsersRepo repositories.CreateUserRepository,
	getUsersRepo repositories.GetUserRepository,
	updateUserRepo repositories.UpdateUserRepository) *UsersUsecase {
	return &UsersUsecase{
		CreateUserUsecase: NewCreateUser(createUsersRepo),
		GetUsersUsecase:   NewGetUsersUsecase(getUsersRepo),
		UpdateUserUsecase: NewUpdateUser(updateUserRepo)}
}
