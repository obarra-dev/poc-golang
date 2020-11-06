package main

import (
	"log"
	"net/http"

	"poc-golang/payment-service/webservices/cbu"
	"poc-golang/payment-service/webservices/product"
)

const basePath = "/api"

func main() {
	cbu.SetupRoutes(basePath)
	product.SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
