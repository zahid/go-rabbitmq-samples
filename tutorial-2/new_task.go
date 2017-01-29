package main

import (
	"fmt"
	"log"
	"os"

	"strings"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatal("wha")
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {

	/* Dial a connection to the RabbitMQ server */

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	/* Channel opens a unique, concurrent server channel to process the bulk of AMQP messages.
	   Any error from methods on this receiver will render the receiver invalid and a new Channel
	   should be opened. */

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	/* QueueDeclare declares a queue to hold messages and deliver to consumers.
	   Declaring creates a queue if it doesn't already exist, or ensures that an existing queue
	   matches the same parameters. */

	q, err := ch.QueueDeclare(
		"task_queue", // name of queue
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	/* Publish sends a Publishing from the client to an exchange on the server. */

	body := bodyFrom(os.Args)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args)) < 2 || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
