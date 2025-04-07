package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reservations-service/internal/domain/models"
)

func GetUserByID(userID int) (*models.User, error) {
	resp, err := http.Get(fmt.Sprintf("http://users-service:8081/users?id=%d", userID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user models.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
