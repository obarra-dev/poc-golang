package main

import (
	"log"
	"net/http"

	"poc-golang/payment-service/webservices/cbu"
	"poc-golang/payment-service/webservices/database"
	"poc-golang/payment-service/webservices/product"
	"poc-golang/payment-service/webservices/report"
)

const basePathAPI = "/api"
const basePathWEB = "/payment-web"

func main() {
	database.SetupDatabase()
	cbu.SetupRoutes(basePathAPI)
	product.SetupRoutes(basePathAPI)
	report.SetupRoutes(basePathWEB)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
