package users

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/users/entities"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/users/repositories"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/database"
)

type UserRepositoryUpdate struct {
	database database.DatabaseGateway
}

func NewUsersRepositoryUpdate(db database.DatabaseGateway) repositories.UpdateUserRepository {
	return &UserRepositoryUpdate{database: db}
}

func (repository *UserRepositoryUpdate) UpdateUser(user entities.Users) error {
	result := repository.database.DB().Model(&entities.Users{}).Where("id = ?", user.ID).Updates(user)
	return result.Error
}
