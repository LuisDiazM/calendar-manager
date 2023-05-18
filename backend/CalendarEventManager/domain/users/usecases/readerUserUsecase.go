package usecases

import (
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/users/entities"
	"github.com/LuisDiazM/calendar-manager/calendar-event-manager/domain/users/repositories"
)

type GetUsersUsecase struct {
	getUserRepo repositories.GetUserRepository
}

func NewGetUsersUsecase(repo repositories.GetUserRepository) *GetUsersUsecase {
	return &GetUsersUsecase{
		getUserRepo: repo,
	}
}

func (usecase *GetUsersUsecase) FindUserById(id string) (entities.Users, error) {
	user, err := usecase.getUserRepo.GetUserById(id)
	if user == nil || err != nil {
		return entities.Users{}, err
	} else {
		return *user, nil
	}
}
