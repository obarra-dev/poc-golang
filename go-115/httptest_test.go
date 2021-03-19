package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServerReturnImmediatelyTimeOut(t *testing.T) {
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusGatewayTimeout)
	}))

	url := backend.URL
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Error("Request error", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error("Response error", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusGatewayTimeout {
		t.Error("Error", resp.StatusCode)
	}
}

func TestServerReturnTimeOutWithHandler(t *testing.T) {
	handlerFuncSleep100ms := http.HandlerFunc(handlerWithSleep100ms)

	//time out when  Any value < 100 ms
	backend := httptest.NewServer(http.TimeoutHandler(handlerFuncSleep100ms, 20*time.Millisecond, "server timeout"))

	url := backend.URL
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Error("Request error", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error("Response error", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Error("Error", resp.StatusCode, resp.Body)
	}
}

func TestServerReturn200(t *testing.T) {
	t.Skip()
	handlerFuncSleep100ms := http.HandlerFunc(handlerWithSleep100ms)

	//time out when  Any value < 100 ms
	backend := httptest.NewServer(http.TimeoutHandler(handlerFuncSleep100ms, 105*time.Millisecond, "server timeout"))

	url := backend.URL
	fmt.Println(url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal("Request error", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal("Response error", err)
		return
	}

	defer resp.Body.Close()

	//TODO a veces tira  http: superfluous response.WriteHeader call from poc-golang/go-115.handlerWithSleep100ms 
	if resp.StatusCode != http.StatusOK {
		t.Fatal("Error", resp.StatusCode, resp.Body)
	}
}

func TestClientSetTimeOut(t *testing.T) {
	handlerFuncSleep100ms := http.HandlerFunc(handlerWithSleep100ms)

	//time out when  Any value < 100 ms
	backend := httptest.NewServer(handlerFuncSleep100ms)

	url := backend.URL
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Error("Request error", err)
		return
	}

	client := http.Client{Timeout: 200 * time.Millisecond}

	resp, err := client.Do(req)
	if err != nil {
		t.Error("Response error", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Error("Error", resp.StatusCode, resp.Body)
	}
}

func handlerWithSleep100ms(w http.ResponseWriter, r *http.Request) {
	time.Sleep(100 * time.Millisecond)
	d := map[string]interface{}{
		"id":    "4040",
		"scope": "test-scope",
	}

	b, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, string(b))
	w.WriteHeader(http.StatusOK)
}
