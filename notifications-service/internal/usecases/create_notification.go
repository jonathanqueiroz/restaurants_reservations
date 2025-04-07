package usecases

import (
	"notifications-service/internal/domain/models"
	"notifications-service/internal/domain/repository"
)

type CreateNotificationUseCase interface {
	Execute(notification *models.Notification) error
}

type createNotificationUseCase struct {
	notificationRepository repository.NotificationRepository
}

func NewCreateNotificationUseCase(notificationRepository repository.NotificationRepository) CreateNotificationUseCase {
	return &createNotificationUseCase{
		notificationRepository: notificationRepository,
	}
}

func (uc *createNotificationUseCase) Execute(notification *models.Notification) error {
	return uc.notificationRepository.CreateNotification(notification)
}
