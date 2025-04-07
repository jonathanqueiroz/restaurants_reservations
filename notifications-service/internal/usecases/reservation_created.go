package usecases

import (
	"encoding/json"
	"log"
	"notifications-service/internal/domain/models"
	"notifications-service/internal/domain/repository"
	"strconv"
)

type ProcessReservationCreatedUseCase struct {
	NotificationRepo repository.NotificationRepository
}

func NewProcessReservationCreatedUseCase(repository repository.NotificationRepository) *ProcessReservationCreatedUseCase {
	return &ProcessReservationCreatedUseCase{NotificationRepo: repository}
}

func (uc *ProcessReservationCreatedUseCase) Execute(message []byte) error {
	var reservation models.Reservation
	if err := json.Unmarshal(message, &reservation); err != nil {
		return err
	}

	notification := models.Notification{
		Message: "Reservation created on Restaurant ID " + strconv.Itoa(reservation.RestaurantID),
	}

	if err := uc.NotificationRepo.CreateNotification(&notification); err != nil {
		return err
	}

	log.Printf("Notification created for user %d: %s", reservation.UserID, notification.Message)
	return nil
}
