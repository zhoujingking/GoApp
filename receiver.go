package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func failOnError(err error, message string) {
	if err != nil {
		log.Panicf("%s, %s", message, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "failed to connect to RabbitMq")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	failOnError(err, "failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	messageChan, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to register a consumer")
	forever := make(chan int)

	go func() {
		for msg := range messageChan {
			log.Printf("received a message %s", msg.Body)
			d := time.Duration(10)
			time.Sleep(d * time.Second)
			msg.Ack(false)
		}
	}()

	<-forever
}
