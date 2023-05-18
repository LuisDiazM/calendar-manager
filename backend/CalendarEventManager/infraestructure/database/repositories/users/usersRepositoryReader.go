package users

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/usecases/users/entities"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/usecases/users/repositories"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/database"
)

type UserRepositoryReads struct {
	database database.DatabaseGateway
}

func NewUsersRepositoryReader(db database.DatabaseGateway) repositories.GetUserRepository {
	return &UserRepositoryReads{
		database: db,
	}
}

func (repository *UserRepositoryReads) GetUserById(id string) (*entities.Users, error) {
	var users entities.Users
	result := repository.database.DB().First(&users, "id = ?", id)
	err := result.Error
	if err != nil {
		return nil, err
	} else {
		return &users, nil
	}
}
