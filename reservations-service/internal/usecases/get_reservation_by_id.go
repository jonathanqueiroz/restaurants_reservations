package usecases

import (
	"reservations-service/internal/domain/models"
	"reservations-service/internal/domain/repository"
)

type GetReservationByIDUseCase interface {
	Execute(id int) (*models.Reservation, error)
}

type getReservationByIDUseCaseImpl struct {
	reservationRepository repository.ReservationRepository
}

func NewGetReservationByIDUseCase(reservationRepository repository.ReservationRepository) GetReservationByIDUseCase {
	return &getReservationByIDUseCaseImpl{
		reservationRepository: reservationRepository,
	}
}

func (uc *getReservationByIDUseCaseImpl) Execute(id int) (*models.Reservation, error) {
	return uc.reservationRepository.GetReservationByID(id)
}
