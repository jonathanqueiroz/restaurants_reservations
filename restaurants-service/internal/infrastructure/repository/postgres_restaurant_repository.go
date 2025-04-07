package repository

import (
	"database/sql"
	"restaurants-service/internal/domain/models"
)

type PostgresRestaurantRepository struct {
	DB *sql.DB
}

func NewPostgresRestaurantRepository(db *sql.DB) *PostgresRestaurantRepository {
	return &PostgresRestaurantRepository{DB: db}
}

func (r *PostgresRestaurantRepository) CreateRestaurant(restaurant *models.Restaurant) error {
	query := "INSERT INTO restaurants (name, address) VALUES ($1, $2) RETURNING id"
	err := r.DB.QueryRow(query, restaurant.Name, restaurant.Address).Scan(&restaurant.ID)
	return err
}

func (r *PostgresRestaurantRepository) GetRestaurantByID(id int) (*models.Restaurant, error) {
	query := "SELECT id, name, address FROM restaurants WHERE id = $1"
	row := r.DB.QueryRow(query, id)

	var restaurant models.Restaurant
	err := row.Scan(&restaurant.ID, &restaurant.Name, &restaurant.Address)
	if err != nil {
		return nil, err
	}
	return &restaurant, nil
}

func (r *PostgresRestaurantRepository) GetAllRestaurants() ([]*models.Restaurant, error) {
	query := "SELECT id, name, address FROM restaurants"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var restaurants []*models.Restaurant
	for rows.Next() {
		var restaurant models.Restaurant
		if err := rows.Scan(&restaurant.ID, &restaurant.Name, &restaurant.Address); err != nil {
			return nil, err
		}
		restaurants = append(restaurants, &restaurant)
	}

	if restaurants == nil {
		return []*models.Restaurant{}, nil
	}

	return restaurants, nil
}
