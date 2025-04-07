package repository

import "notifications-service/internal/domain/models"

type NotificationRepository interface {
	CreateNotification(notification *models.Notification) error
	GetNotificationByID(id int) (*models.Notification, error)
	GetAllNotifications() ([]*models.Notification, error)
}
