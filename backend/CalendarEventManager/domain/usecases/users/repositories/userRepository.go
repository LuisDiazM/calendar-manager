package repositories

import "github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/usecases/users/entities"

type CreateUserRepository interface {
	CreateUser(user entities.Users) error
}
