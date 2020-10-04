package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

type myHandler struct {
	message string
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(*h)
	w.Write(body)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/handleFunc", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>"+html.EscapeString(r.URL.Path)+"</h1>")
	})

	mux.HandleFunc("/handleFunction", handleFunction)

	myHanlder := &myHandler{"Hi World with handler"}

	mux.Handle("/handle", myHanlder)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
}

func handleFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>"+html.EscapeString(r.URL.Path)+"</h1>")
}
