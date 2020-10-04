package main

import (
	"io"
	"log"
	"net"
	"time"
)

func mainNet() {
	listener, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatal("Error in creating listener")
	}

	log.Println("Server UP")

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println(err, "Error connecting")
			continue
		}

		go manageConnection(connection)
	}
}

func manageConnection(connection net.Conn) {
	defer connection.Close()
	log.Println("Got Connection", connection)
	for {
		_, err := io.WriteString(connection, time.Now().Format(time.RFC3339))
		if err != nil {
			log.Fatal(err, "Error sending data to client")
			return
		}

		time.Sleep(1 * time.Second)
	}
}
