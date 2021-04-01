package main

import (
	"log"
	"net/http"

	websockets "poc-golang/payment-service/webservices/websocket"

	"poc-golang/payment-service/webservices/cbu"
	"poc-golang/payment-service/webservices/database"
	"poc-golang/payment-service/webservices/product"
	"poc-golang/payment-service/webservices/report"
	"poc-golang/payment-service/webservices/user"
)

const basePathAPI = "/api"
const basePathWEB = "/payment-web"
const basePathWS = "/ws"

func main() {
	database.SetupDatabase()

	cbu.SetupRoutes(basePathAPI)
	product.SetupRoutes(basePathAPI)
	user.SetupRoutes(basePathAPI)

	report.SetupRoutes(basePathWEB)

	websockets.SetupRoutes(basePathWS)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
