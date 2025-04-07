package usecases

import (
	"restaurants-service/internal/domain/models"
	"restaurants-service/internal/domain/repository"
)

type GetAllRestaurantsUseCase interface {
	Execute() ([]*models.Restaurant, error)
}

type getAllRestaurantsUseCaseImpl struct {
	restaurantRepository repository.RestaurantRepository
}

func NewGetAllRestaurantsUseCase(restaurantRepository repository.RestaurantRepository) GetAllRestaurantsUseCase {
	return &getAllRestaurantsUseCaseImpl{
		restaurantRepository: restaurantRepository,
	}
}

func (uc *getAllRestaurantsUseCaseImpl) Execute() ([]*models.Restaurant, error) {
	return uc.restaurantRepository.GetAllRestaurants()
}
