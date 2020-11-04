package main

import (
	"log"
	"net/http"

	"poc-golang/payment-service/webservices/cbu"
)

const basePath = "/api"

func main() {
	cbu.SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
