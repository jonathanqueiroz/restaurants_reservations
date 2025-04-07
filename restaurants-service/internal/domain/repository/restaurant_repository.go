package repository

import "restaurants-service/internal/domain/models"

type RestaurantRepository interface {
	CreateRestaurant(restaurant *models.Restaurant) error
	GetRestaurantByID(id int) (*models.Restaurant, error)
	GetAllRestaurants() ([]*models.Restaurant, error)
}
