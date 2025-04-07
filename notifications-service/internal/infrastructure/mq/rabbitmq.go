package mq

import (
	"encoding/json"
	"fmt"
	"log"
	"notifications-service/internal/domain/models"
	"notifications-service/internal/domain/repository"
	"os"

	"github.com/streadway/amqp"
)

var Channel *amqp.Channel

func InitRabbitMQ() {
	rabbitmqHost := os.Getenv("RABBITMQ_HOST")
	rabbitmqPort := os.Getenv("RABBITMQ_PORT")
	rabbitmqURL := fmt.Sprintf("amqp://guest:guest@%s:%s/", rabbitmqHost, rabbitmqPort)

	var err error
	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		log.Fatal(err)
	}

	Channel, err = conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
}

func ConsumeMessages(notificationRepo repository.NotificationRepository) {
	// Declaração da fila "user_created"
	userQueue, err := Channel.QueueDeclare(
		"user_created", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	// Consome mensagens da fila "user_created"
	userMsgs, err := Channel.Consume(
		userQueue.Name, // queue
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	if err != nil {
		log.Fatal(err)
	}

	// Processa mensagens da fila "user_created"
	go func() {
		for d := range userMsgs {
			var user models.User
			if err := json.Unmarshal(d.Body, &user); err != nil {
				log.Printf("Error decoding JSON: %s", err)
				continue
			}
			notification := models.Notification{
				UserID:  user.ID,
				Message: "Welcome, " + user.Name + "!",
			}
			if err := notificationRepo.CreateNotification(&notification); err != nil {
				log.Printf("Error creating notification: %s", err)
			} else {
				log.Printf("Notification created for user %d: %s", user.ID, notification.Message)
			}
		}
	}()

	// Declaração da fila "reservation_created"
	reservationQueue, err := Channel.QueueDeclare(
		"reservation_created", // name
		false,                 // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	// Consome mensagens da fila "reservation_created"
	reservationMsgs, err := Channel.Consume(
		reservationQueue.Name, // queue
		"",                    // consumer
		true,                  // auto-ack
		false,                 // exclusive
		false,                 // no-local
		false,                 // no-wait
		nil,                   // args
	)
	if err != nil {
		log.Fatal(err)
	}

	// Processa mensagens da fila "reservation_created"
	go func() {
		for d := range reservationMsgs {
			var reservation models.Reservation
			if err := json.Unmarshal(d.Body, &reservation); err != nil {
				log.Printf("Error decoding JSON: %s", err)
				continue
			}
			notification := models.Notification{
				UserID:  reservation.UserID,
				Message: "Reservation created on Restaurant ID " + fmt.Sprint(reservation.RestaurantID),
			}
			if err := notificationRepo.CreateNotification(&notification); err != nil {
				log.Printf("Error creating notification: %s", err)
			} else {
				log.Printf("Notification created for user %d: %s", reservation.UserID, notification.Message)
			}
		}
	}()
}
