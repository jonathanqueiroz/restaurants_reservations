package main

import (
	"log"
	"net/http"
	"reservations-service/internal/infrastructure/db"
	"reservations-service/internal/infrastructure/mq"
	"reservations-service/internal/infrastructure/repository"
	"reservations-service/internal/interfaces/handlers"
	"reservations-service/internal/usecases"
)

func main() {
	// Inicializa o banco de dados
	db.InitDB()
	// Inicializa o RabbitMQ
	mq.InitRabbitMQ()

	// Inicializa o repositório e os casos de uso
	reservationRepository := repository.NewPostgresReservationRepository(db.DB)
	createReservationUseCase := usecases.NewCreateReservationUseCase(reservationRepository)
	getReservationByIDUseCase := usecases.NewGetReservationByIDUseCase(reservationRepository)
	getAllReservationsUseCase := usecases.NewGetAllReservationsUseCase(reservationRepository)

	reservationHandler := handlers.NewReservationHandler(createReservationUseCase, getReservationByIDUseCase, getAllReservationsUseCase)
	http.HandleFunc("/reservations", reservationHandler.ServeHTTP)

	// Inicia o servidor HTTP
	log.Println("Starting server on :8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
