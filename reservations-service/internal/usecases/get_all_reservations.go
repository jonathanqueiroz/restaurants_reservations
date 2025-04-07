package usecases

import (
	"reservations-service/internal/domain/models"
	"reservations-service/internal/domain/repository"
)

type GetAllReservationsUseCase interface {
	Execute() ([]*models.Reservation, error)
}

type getAllReservationsUseCaseImpl struct {
	reservationRepository repository.ReservationRepository
}

func NewGetAllReservationsUseCase(reservationRepository repository.ReservationRepository) GetAllReservationsUseCase {
	return &getAllReservationsUseCaseImpl{
		reservationRepository: reservationRepository,
	}
}

func (uc *getAllReservationsUseCaseImpl) Execute() ([]*models.Reservation, error) {
	return uc.reservationRepository.GetAllReservations()
}
