package users

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/users/entities"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/users/repositories"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/infraestructure/database"
)

type UserRepository struct {
	database database.DatabaseGateway
}

func NewUsersRepository(db database.DatabaseGateway) repositories.CreateUserRepository {
	return &UserRepository{database: db}
}

func (repository *UserRepository) CreateUser(user entities.Users) error {
	result := repository.database.DB().Create(&user)
	return result.Error
}
