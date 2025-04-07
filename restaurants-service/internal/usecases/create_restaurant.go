package usecases

import (
	"restaurants-service/internal/domain/models"
	"restaurants-service/internal/domain/repository"
)

type CreateRestaurantUseCase interface {
	Execute(restaurant *models.Restaurant) error
}

type createRestaurantUseCaseImpl struct {
	restaurantRepository repository.RestaurantRepository
}

func NewCreateRestaurantUseCase(restaurantRepository repository.RestaurantRepository) CreateRestaurantUseCase {
	return &createRestaurantUseCaseImpl{
		restaurantRepository: restaurantRepository,
	}
}

func (uc *createRestaurantUseCaseImpl) Execute(restaurant *models.Restaurant) error {
	return uc.restaurantRepository.CreateRestaurant(restaurant)
}
