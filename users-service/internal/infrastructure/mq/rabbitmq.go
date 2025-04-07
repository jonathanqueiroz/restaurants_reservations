package mq

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

var Conn *amqp.Connection
var Channel *amqp.Channel

func InitRabbitMQ() {
	rabbitmqHost := os.Getenv("RABBITMQ_HOST")
	rabbitmqPort := os.Getenv("RABBITMQ_PORT")
	rabbitmqURL := fmt.Sprintf("amqp://guest:guest@%s:%s/", rabbitmqHost, rabbitmqPort)

	var err error
	Conn, err = amqp.Dial(rabbitmqURL)
	if err != nil {
		log.Fatal(err)
	}

	Channel, err = Conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
}

func PublishMessage(queueName string, message []byte) error {
	err := Channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		})
	if err != nil {
		return err
	}
	log.Printf("Message published to queue %s: %s", queueName, message)
	return nil
}
