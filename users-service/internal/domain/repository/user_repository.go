package repository

import "users-service/internal/domain/models"

type UserRepository interface {
	CreateUser(user *models.User) (int, error)
	GetUserByID(id int) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
}
