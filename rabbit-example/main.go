package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// to create the rabbit-server
// docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management

func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672")

	failOnError(err, "Failed to connect to RabbitMQ")
	defer connection.Close()

	ch, err := connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	args := make(amqp.Table)
	args["x-dead-letter-exchange"] = ""
	args["x-dead-letter-routing-key"] = "ANALYSIS_EVENT_DEAD_LETTER"
	args["x-message-ttl"] = int32(10800000)

	// it is needed just one time
	queue, err := ch.QueueDeclare(
		"ANALYSIS_EVENT",
		true,
		false,
		false,
		false,
		args,
	)
	failOnError(err, "Failed to declare a Queue")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	m := make(map[string]string)
	m["k1"] = "omar"

	jobEvent := &JobEvent{Identifier: "Frank", JobEventResult: "SUCCESS", Labels: m}
	b, err := json.Marshal(jobEvent)
	failOnError(err, "Failed to marshal ")

	err = ch.PublishWithContext(ctx,
		"", // exchange
		"ANALYSIS_EVENT",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		})

	failOnError(err, "Failed to publish a message")

	fmt.Println("Queue status:", queue)
	log.Printf(" [x] Sent %s", b)
}

type JobEvent struct {
	Identifier     string            `json:"identifier"`
	JobEventResult string            `json:"jobEventResult"`
	Labels         map[string]string `json:"labels"`
	Failure        *JobEventFailure  `json:"failure"`
}

type JobEventFailure struct {
	Reason  *string `json:"reason"`
	Message *string `json:"message"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
