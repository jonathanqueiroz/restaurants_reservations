package usecases

import (
	"encoding/json"
	"log"
	"notifications-service/internal/domain/models"
	"notifications-service/internal/domain/repository"
)

type ProcessUserCreatedUseCase struct {
	NotificationRepository repository.NotificationRepository
}

func NewProcessUserCreatedUseCase(notificationRepository repository.NotificationRepository) *ProcessUserCreatedUseCase {
	return &ProcessUserCreatedUseCase{
		NotificationRepository: notificationRepository,
	}
}

func (uc *ProcessUserCreatedUseCase) Execute(message []byte) error {
	var user models.User
	if err := json.Unmarshal(message, &user); err != nil {
		return err
	}

	notification := models.Notification{
		UserID:  user.ID,
		Message: "Welcome, " + user.Name + "!",
	}

	if err := uc.NotificationRepository.CreateNotification(&notification); err != nil {
		return err
	}

	log.Printf("Notification created for user %d: %s", user.ID, notification.Message)
	return nil
}
