package repository

import (
	"database/sql"
	"reservations-service/internal/domain/models"
)

type PostgresReservationRepository struct {
	DB *sql.DB
}

func NewPostgresReservationRepository(db *sql.DB) *PostgresReservationRepository {
	return &PostgresReservationRepository{
		DB: db,
	}
}

func (r *PostgresReservationRepository) CreateReservation(reservation *models.Reservation) error {
	query := "INSERT INTO reservations (user_id, restaurant_id, date, time) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.DB.QueryRow(query, reservation.UserID, reservation.RestaurantID, reservation.Date, reservation.Time).Scan(&reservation.ID)
	return err
}

func (r *PostgresReservationRepository) GetReservationByID(id int) (*models.Reservation, error) {
	query := "SELECT id, user_id, restaurant_id, date, time FROM reservations WHERE id = $1"
	row := r.DB.QueryRow(query, id)

	var reservation models.Reservation
	err := row.Scan(&reservation.ID, &reservation.UserID, &reservation.RestaurantID, &reservation.Date, &reservation.Time)
	if err != nil {
		return nil, err
	}
	return &reservation, nil
}

func (r *PostgresReservationRepository) GetAllReservations() ([]*models.Reservation, error) {
	query := "SELECT id, user_id, restaurant_id, date, time FROM reservations"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []*models.Reservation
	for rows.Next() {
		var reservation models.Reservation
		if err := rows.Scan(&reservation.ID, &reservation.UserID, &reservation.RestaurantID, &reservation.Date, &reservation.Time); err != nil {
			return nil, err
		}
		reservations = append(reservations, &reservation)
	}

	if reservations == nil {
		return []*models.Reservation{}, nil
	}

	return reservations, nil
}
