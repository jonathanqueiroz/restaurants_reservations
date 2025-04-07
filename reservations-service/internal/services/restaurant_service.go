package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reservations-service/internal/domain/models"
)

func GetRestaurantByID(restaurantID int) (*models.Restaurant, error) {
	resp, err := http.Get(fmt.Sprintf("http://restaurants-service:8082/restaurants?id=%d", restaurantID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var restaurant models.Restaurant
	if err := json.NewDecoder(resp.Body).Decode(&restaurant); err != nil {
		return nil, err
	}
	return &restaurant, nil
}
