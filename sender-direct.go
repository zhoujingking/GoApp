package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"time"
)

func failOnError(err error, message string) {
	if err != nil {
		log.Panicf("%s, %s", message, err)
	}
}

func main() {
	if len(os.Args) != 3 {
		log.Panicf("usage: sender-direct type message")
		return
	}
	routeKey := os.Args[1]
	message := os.Args[2]

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare("direct-demo", "direct", true, false, false, false, nil)
	failOnError(err, "failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"direct-demo",
		routeKey, // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	failOnError(err, "failed to publish a message")
	log.Printf("message sent successfully")
}
