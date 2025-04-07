package repository

import (
	"database/sql"
	"fmt"
	"users-service/internal/domain/models"
)

type PostgresUserRepository struct {
	DB *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		DB: db,
	}
}

func (r *PostgresUserRepository) CreateUser(user *models.User) (int, error) {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	var id int

	err := r.DB.QueryRow(query, user.Name, user.Email).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error creating user: %w", err)
	}

	user.ID = id

	return id, nil
}

func (r *PostgresUserRepository) GetUserByID(id int) (*models.User, error) {
	query := `SELECT id, name, email FROM users WHERE id = $1`
	row := r.DB.QueryRow(query, id)

	var user models.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error getting user by ID: %w", err)
	}
	return &user, nil
}

func (r *PostgresUserRepository) GetAllUsers() ([]*models.User, error) {
	query := `SELECT id, name, email FROM users`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error getting all users: %w", err)
	}
	defer rows.Close()
	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, fmt.Errorf("error scanning user: %w", err)
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}
	return users, nil
}
