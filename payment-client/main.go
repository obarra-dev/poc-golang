package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"golang.org/x/net/websocket"
)

const address string = "localhost:5000"

type message struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

type product struct {
	ProductID      *int   `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

func main() {
	initWebsocketClient()
}

func initWebsocketClient() {
	fmt.Println("Starting Client")
	ws, err := websocket.Dial(fmt.Sprintf("ws://%s/ws/products", address), "", fmt.Sprintf("http://%s/", address))
	if err != nil {
		fmt.Printf("Dial failed: %s\n", err.Error())
		os.Exit(1)
	}
	incomingProducts := make(chan []product)
	go readClientMessages(ws, incomingProducts)
	i := 0
	for {
		select {
		case <-time.After(time.Duration(2e9)):
			i++
			response := new(message)
			response.Data = "Sent data " + strconv.Itoa(i)
			response.Type = strconv.Itoa(i)
			err = websocket.JSON.Send(ws, response)
			if err != nil {
				fmt.Printf("Send failed: %s\n", err.Error())
				os.Exit(1)
			}
		case message := <-incomingProducts:
			fmt.Println(`Message Received:`, message)

		}
	}
}

func readClientMessages(ws *websocket.Conn, incomingProducts chan []product) {
	for {
		var productList []product
		if err := websocket.JSON.Receive(ws, &productList); err != nil {
			fmt.Printf("Error::: %s\n", err.Error())
			return
		}

		incomingProducts <- productList
	}
}
