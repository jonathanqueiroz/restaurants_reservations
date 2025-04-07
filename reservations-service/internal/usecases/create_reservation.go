package usecases

import (
	"reservations-service/internal/domain/models"
	"reservations-service/internal/domain/repository"
)

type CreateReservationUseCase interface {
	Execute(reservation *models.Reservation) error
}

type createReservationUseCaseImpl struct {
	reservationRepository repository.ReservationRepository
}

func NewCreateReservationUseCase(reservationRepository repository.ReservationRepository) CreateReservationUseCase {
	return &createReservationUseCaseImpl{
		reservationRepository: reservationRepository,
	}
}

func (uc *createReservationUseCaseImpl) Execute(reservation *models.Reservation) error {
	return uc.reservationRepository.CreateReservation(reservation)
}
