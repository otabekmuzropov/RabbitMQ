package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {

	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		log.Panic("Failed connect to RabbitMQ", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Panic("Failed to create new channel", err)
	}

	defer ch.Close()

	msg := amqp.Publishing{
		Body:[]byte("my first"),
	}

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


	err = ch.Publish(
		"test",
		"test.create",
		false,
		false,
		msg,
		)

	if err != nil {
		log.Panic("Error while publishing", err)
	}


}