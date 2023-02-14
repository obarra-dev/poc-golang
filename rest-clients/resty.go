package main

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"host":"github.com","repoOwner":"obarra-dev","repoName":"minesweeper-API","branch":"add-errors","generatePrToFixIssues":false}`).
		Post("https://reqres.in/api/users")

	failOnError(err, "Error doing POST")
	fmt.Println("Resp: ", resp)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s, %s", msg, err, err.Error())
	}
}
