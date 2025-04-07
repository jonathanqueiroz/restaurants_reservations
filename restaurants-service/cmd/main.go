package main

import (
	"log"
	"net/http"
	"restaurants-service/internal/infrastructure/db"
	"restaurants-service/internal/infrastructure/repository"
	"restaurants-service/internal/interfaces/handlers"
	"restaurants-service/internal/usecases"
)

func main() {
	// Inicializa o banco de dados
	db.InitDB()

	// Inicializa o handler com a instância do banco de dados
	restaurantRepository := repository.NewPostgresRestaurantRepository(db.DB)
	createRestaurantUseCase := usecases.NewCreateRestaurantUseCase(restaurantRepository)
	getRestaurantByIDUseCase := usecases.NewGetRestaurantByIDUseCase(restaurantRepository)
	getAllRestaurantsUseCase := usecases.NewGetAllRestaurantsUseCase(restaurantRepository)

	restaurantHandler := handlers.NewRestaurantHandler(createRestaurantUseCase, getRestaurantByIDUseCase, getAllRestaurantsUseCase)
	http.HandleFunc("/restaurants", restaurantHandler.ServeHTTP)

	// Inicia o servidor HTTP
	log.Println("Starting server on :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
