package usecases

import (
	"users-service/internal/domain/models"
	"users-service/internal/domain/repository"
)

type GetUserByIDUseCase interface {
	Execute(id int) (*models.User, error)
}

type getUserByIDUseCaseImpl struct {
	UserRepository repository.UserRepository
}

func NewGetUserByIDUseCase(userRepository repository.UserRepository) GetUserByIDUseCase {
	return &getUserByIDUseCaseImpl{
		UserRepository: userRepository,
	}
}

func (uc *getUserByIDUseCaseImpl) Execute(id int) (*models.User, error) {
	return uc.UserRepository.GetUserByID(id)
}
