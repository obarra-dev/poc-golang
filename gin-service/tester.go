package main

import (
	fmt "fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func mainxx() {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/videos", nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Authorization", "Basic b21hcjoxMjM=")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))
}
