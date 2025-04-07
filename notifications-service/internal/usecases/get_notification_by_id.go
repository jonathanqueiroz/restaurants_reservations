package usecases

import (
	"notifications-service/internal/domain/models"
	"notifications-service/internal/domain/repository"
)

type GetNotificationByIDUseCase interface {
	Execute(id int) (*models.Notification, error)
}

type getNotificationByIDUseCase struct {
	notificationRepository repository.NotificationRepository
}

func NewGetNotificationByIDUseCase(notificationRepository repository.NotificationRepository) GetNotificationByIDUseCase {
	return &getNotificationByIDUseCase{
		notificationRepository: notificationRepository,
	}
}

func (uc *getNotificationByIDUseCase) Execute(id int) (*models.Notification, error) {
	return uc.notificationRepository.GetNotificationByID(id)
}
