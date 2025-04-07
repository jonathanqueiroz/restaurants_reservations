package models

type Reservation struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	RestaurantID int    `json:"restaurant_id"`
	Date         string `json:"date"`
	Time         string `json:"time"`
}
