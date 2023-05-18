package repositories

import "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/usecases/users/entities"

type CreateUserRepository interface {
	CreateUser(user entities.Users) error
}

type GetUserRepository interface {
	GetUserById(id string) (*entities.Users, error)
}

type UpdateUserRepository interface {
	UpdateUser(user entities.Users) error
}
