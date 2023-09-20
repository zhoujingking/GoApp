package main

import (
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
	if len(os.Args) != 2 {
		log.Panicf("usage: sender-direct type")
		return
	}
	routeKey := os.Args[1]

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "failed to connect to RabbitMq")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare("direct-demo", "direct", true, false, false, false, nil)
	failOnError(err, "failed to declare an exchange")

	// queue name is autogenerated for us
	q, err := ch.QueueDeclare("", false, false, true, false, nil)
	failOnError(err, "failed to declare a queue")

	err = ch.QueueBind(q.Name, routeKey, "direct-demo", false, nil)
	failOnError(err, "failed to bind a queue")

	err = ch.QueueBind(q.Name, "all", "direct-demo", false, nil)
	failOnError(err, "failed to bind a queue with key all")

	messageChan, err := ch.Consume(
		q.Name,
		"",
		true,
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
		}
	}()

	<-forever
}
