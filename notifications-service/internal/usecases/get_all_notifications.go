package usecases

import (
	"notifications-service/internal/domain/models"
	"notifications-service/internal/domain/repository"
)

type GetAllNotificationsUseCase interface {
	Execute() ([]*models.Notification, error)
}

type getAllNotificationsUseCase struct {
	notificationRepository repository.NotificationRepository
}

func NewGetAllNotificationsUseCase(notificationRepository repository.NotificationRepository) GetAllNotificationsUseCase {
	return &getAllNotificationsUseCase{
		notificationRepository: notificationRepository,
	}
}

func (uc *getAllNotificationsUseCase) Execute() ([]*models.Notification, error) {
	return uc.notificationRepository.GetAllNotifications()
}
