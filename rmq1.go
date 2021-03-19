package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main()  {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	defer conn.Close()

	if err != nil {
		log.Panic("Failed to connect to RabbitMQ", err)
	}

	ch, err := conn.Channel()

	if err != nil {
		log.Panic("Failed to create a new channel", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"test",
		amqp.ExchangeTopic,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Panic("Error while declaring exchange")
	}

	queue, err := ch.QueueDeclare(
		"test.test1",
		false,
		false,
		true,
		false,
		nil,
	)

	if err != nil {
		log.Panic("Error to declare queue")
	}

	err = ch.QueueBind(
		queue.Name,
		"test.#",
		"test",
		false,
		nil,
	)
	if err != nil {
		log.Panic("Error while binding to exchange")
	}

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	<-forever
}
