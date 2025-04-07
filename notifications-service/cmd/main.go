package main

import (
	"log"
	"net/http"
	"notifications-service/internal/infrastructure/db"
	"notifications-service/internal/infrastructure/mq"
	"notifications-service/internal/infrastructure/repository"
	"notifications-service/internal/interfaces/handlers"
	"notifications-service/internal/usecases"
)

func main() {
	// Inicializa o banco de dados
	db.InitDB()

	// Inicializa o repositório
	notificationRepository := repository.NewPostgresNotificationRepository(db.DB)

	// Inicializa o RabbitMQ
	mq.InitRabbitMQ()
	mq.ConsumeMessages(notificationRepository)

	// Inicializa o repositório de mensagens
	createNotificationUseCase := usecases.NewCreateNotificationUseCase(notificationRepository)
	getNotificationByIDUseCase := usecases.NewGetNotificationByIDUseCase(notificationRepository)
	getAllNotificationsUseCase := usecases.NewGetAllNotificationsUseCase(notificationRepository)

	// Cria o handler para as notificações
	notificationHandler := handlers.NewNotificationHandler(createNotificationUseCase, getNotificationByIDUseCase, getAllNotificationsUseCase)

	// Define o handler para a rota /notifications
	http.HandleFunc("/notifications", notificationHandler.ServeHTTP)

	// Inicia o servidor HTTP
	log.Println("Starting server on :8084")
	log.Fatal(http.ListenAndServe(":8084", nil))
}
