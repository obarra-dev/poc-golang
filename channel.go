package main

import (
	"fmt"
	"time"
)

func executeChanelBuffered() {
	messages := make(chan string, 5)
	for i := 0; i < 5; i++ {
		go sendMessage(messages, i)
	}

	showMessage(messages)
}

func sendMessage(messages chan<- string, goRutine int) {
	counter := 0
	for {
		messages <- "GR-" + fmt.Sprint(goRutine) + "You are go developer " + fmt.Sprint(counter)
		fmt.Println("GR: ", goRutine, " - sending message number: ", counter)
		counter++
	}
}

func showMessage(messages <-chan string) {
	for message := range messages {
		fmt.Println("Got: ", message)
		time.Sleep(1 * time.Second)
	}
}

func executeRangeAndUndirection() {
	number := make(chan int)
	pow := make(chan int)

	go generateNumbers(number)
	go powTwo(pow, number)

	showNumbers(pow)
}

func generateNumbers(out chan<- int) {
	for x := 1; x < 6; x++ {
		out <- x
	}
	close(out)
}

func powTwo(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func showNumbers(in <-chan int) {
	for x := range in {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}

func executeChannelsExample() {

	channel := make(chan string)

	go echoPingServer(channel)
	go sendPingClient(channel)

	var input string
	fmt.Scanln(&input)
	fmt.Println("End  example")
}

func echoPingServer(channel chan string) {
	var counter int
	for {
		counter++
		value, ok := <-channel
		if !ok {
			break
		}
		fmt.Println(value, " ", counter)
		time.Sleep(time.Second * 1)
	}
}

func sendPingClient(channel chan string) {
	limit := 0
	for {
		channel <- "ping"
		limit++
		if limit > 20 {
			close(channel)
		}
	}
}
