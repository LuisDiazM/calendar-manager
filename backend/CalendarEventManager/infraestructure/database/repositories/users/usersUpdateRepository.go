package users

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/usecases/users/entities"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/usecases/users/repositories"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/database"
)

type UserRepositoryUpdate struct {
	database database.DatabaseGateway
}

func NewUsersRepositoryUpdate(db database.DatabaseGateway) repositories.UpdateUserRepository {
	return &UserRepositoryUpdate{database: db}
}

func (repository *UserRepositoryUpdate) UpdateUser(user entities.Users) error {
	result := repository.database.DB().Model(&entities.Users{}).Where("id = ?", user.Id).Updates(user)
	return result.Error
}
