package main

import (
	"context"
	"log"
	"time"

	"poc-golang/sqs-service/internal/email"
	"poc-golang/sqs-service/internal/pkg/cloud/aws"
)

func main() {
	// Create a cancellable context.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a session instance.
	ses, err := aws.New(aws.Config{
		Address: "http://localhost:4566",
		Region:  "eu-west-1",
		Profile: "localstack",
		ID:      "test",
		Secret:  "test",
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Set queue URL.
	url := "http://localhost:4566/000000000000/welcome-email"

	// Instantiate client.
	client := aws.NewSQS(ses, time.Second*5)

	// Instantiate consumer and start consuming.
	email.NewConsumer(client, email.ConsumerConfig{
		Type:      email.AsyncConsumer,
		QueueURL:  url,
		MaxWorker: 2,
		MaxMsg:    10,
	}).Start(ctx)
}
