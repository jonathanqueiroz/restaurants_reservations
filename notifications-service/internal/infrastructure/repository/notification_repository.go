package repository

import (
	"database/sql"
	"notifications-service/internal/domain/models"
)

type NotificationRepository struct {
	DB *sql.DB
}

func NewPostgresNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{DB: db}
}

func (repo *NotificationRepository) CreateNotification(notification *models.Notification) error {
	query := "INSERT INTO notifications (user_id, message) VALUES ($1, $2) RETURNING id"
	err := repo.DB.QueryRow(query, notification.UserID, notification.Message).Scan(&notification.ID)
	return err
}

func (repo *NotificationRepository) GetNotificationByID(id int) (*models.Notification, error) {
	query := "SELECT id, user_id, message FROM notifications WHERE id = $1"
	row := repo.DB.QueryRow(query, id)

	var notification models.Notification
	err := row.Scan(&notification.ID, &notification.UserID, &notification.Message)
	if err != nil {
		return nil, err
	}
	return &notification, nil
}

func (repo *NotificationRepository) GetAllNotifications() ([]*models.Notification, error) {
	query := "SELECT id, user_id, message FROM notifications"
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []*models.Notification
	for rows.Next() {
		var notification models.Notification
		if err := rows.Scan(&notification.ID, &notification.UserID, &notification.Message); err != nil {
			return nil, err
		}
		notifications = append(notifications, &notification)
	}

	if notifications == nil {
		return []*models.Notification{}, nil
	}

	return notifications, nil
}
