package usecases

import (
	"users-service/internal/domain/models"
	"users-service/internal/domain/repository"
)

type CreateUserUseCase interface {
	Execute(user *models.User) error
}

type createUserUseCaseImpl struct {
	userRepository repository.UserRepository
}

func NewCreateUserUseCase(userRepository repository.UserRepository) CreateUserUseCase {
	return &createUserUseCaseImpl{
		userRepository: userRepository,
	}
}

func (uc *createUserUseCaseImpl) Execute(user *models.User) error {
	userID, err := uc.userRepository.CreateUser(user)
	if err != nil {
		return err
	}

	user.ID = userID
	return nil
}
