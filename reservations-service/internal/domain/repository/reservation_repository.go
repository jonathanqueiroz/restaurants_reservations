package repository

import "reservations-service/internal/domain/models"

type ReservationRepository interface {
	CreateReservation(reservation *models.Reservation) error
	GetReservationByID(id int) (*models.Reservation, error)
	GetAllReservations() ([]*models.Reservation, error)
}
