package usecases

import (
	"users-service/internal/domain/models"
	"users-service/internal/domain/repository"
)

type GetAllUsersUseCase interface {
	Execute() ([]*models.User, error)
}

type getAllUsersUseCaseImpl struct {
	UserRepository repository.UserRepository
}

func NewGetAllUsersUseCase(userRepository repository.UserRepository) GetAllUsersUseCase {
	return &getAllUsersUseCaseImpl{
		UserRepository: userRepository,
	}
}

func (uc *getAllUsersUseCaseImpl) Execute() ([]*models.User, error) {
	return uc.UserRepository.GetAllUsers()
}
