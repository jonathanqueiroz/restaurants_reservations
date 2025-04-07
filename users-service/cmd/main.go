package main

import (
	"log"
	"net/http"
	"users-service/internal/infrastructure/db"
	"users-service/internal/infrastructure/mq"
	"users-service/internal/infrastructure/repository"
	"users-service/internal/interfaces/handlers"
	"users-service/internal/usecases"
)

func main() {
	db.InitDB()
	mq.InitRabbitMQ()

	userRepository := repository.NewPostgresUserRepository(db.DB)
	createUserUseCase := usecases.NewCreateUserUseCase(userRepository)
	getUserByIDUseCase := usecases.NewGetUserByIDUseCase(userRepository)
	getAllUsersUseCase := usecases.NewGetAllUsersUseCase(userRepository)

	userHandler := handlers.NewUserHandler(createUserUseCase, getUserByIDUseCase, getAllUsersUseCase)
	http.HandleFunc("/users", userHandler.ServeHTTP)

	log.Println("Server running on port :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
