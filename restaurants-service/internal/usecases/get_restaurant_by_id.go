package usecases

import (
	"restaurants-service/internal/domain/models"
	"restaurants-service/internal/domain/repository"
)

type GetRestaurantByIDUseCase interface {
	Execute(id int) (*models.Restaurant, error)
}

type getRestaurantByIDUseCaseImpl struct {
	restaurantRepository repository.RestaurantRepository
}

func NewGetRestaurantByIDUseCase(restaurantRepository repository.RestaurantRepository) GetRestaurantByIDUseCase {
	return &getRestaurantByIDUseCaseImpl{
		restaurantRepository: restaurantRepository,
	}
}

func (uc *getRestaurantByIDUseCaseImpl) Execute(id int) (*models.Restaurant, error) {
	return uc.restaurantRepository.GetRestaurantByID(id)
}
