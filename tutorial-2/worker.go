package main

import (
	"fmt"
	"log"

	"bytes"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
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

	/* Consume immediately starts delivering queued messages.
	   Begin receiving on the returned chan Delivery before any other operation on the Connection
	   or Channel. */

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for incomingMessage := range msgs {
			log.Printf("Received a message : %s", incomingMessage.Body)
			numberOfDots := bytes.Count(incomingMessage.Body, []byte("."))
			workTime := time.Duration(numberOfDots)
			time.Sleep(workTime * time.Second)
			log.Printf("Done")
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
