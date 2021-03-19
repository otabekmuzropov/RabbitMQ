package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
)

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1]  == "" {
		s = "hello world!"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")

	if err != nil {
		log.Panic("Failed to connect rabbit", err)
	}

	ch, err := conn.Channel()

	if err != nil {
		log.Panic("Failed to create new channel", err)
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
		log.Panic("Error while declaring exchange", err)
	}

	body := bodyFrom(os.Args)

	err = ch.Publish(
		"test",
		"test.create",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	if err != nil {
		log.Panic(err)
	}

	log.Printf(" [x] Sent %s", []byte(body))


}