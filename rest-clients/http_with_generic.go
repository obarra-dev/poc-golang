package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func httpWithGeneric() {
	ctx := context.Background()
	timeout := 30 * time.Second

	// Get data
	reqContext, _ := context.WithTimeout(ctx, timeout)
	m, err := Get[RequestObj](reqContext, "https://reqres.in/api/users?page=2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m.Data[0])

	// Post data
	user := User{Name: "obarra", Job: "software engineer"}
	addContext, _ := context.WithTimeout(ctx, timeout)
	newUser, err := Post[User](addContext, "https://reqres.in/api/users", user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newUser)
}

type RequestObj struct {
	TotalPage int       `json:"total"`
	Data      []Profile `json:"data"`
}
type Profile struct {
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type User struct {
	Name      string    `json:"name,omitempty"`
	Job       string    `json:"job,omitempty"`
	ID        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func Get[T any](ctx context.Context, url string) (T, error) {
	var m T
	r, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return m, err
	}
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return m, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return m, err
	}
	return parseJSON[T](body)
}
func parseJSON[T any](s []byte) (T, error) {
	var r T
	if err := json.Unmarshal(s, &r); err != nil {
		return r, err
	}
	return r, nil
}

func Post[T any](ctx context.Context, url string, data any) (T, error) {
	var m T
	b, err := toJSON(data)
	if err != nil {
		return m, err
	}
	byteReader := bytes.NewReader(b)
	r, err := http.NewRequestWithContext(ctx, "POST", url, byteReader)
	if err != nil {
		return m, err
	}
	// Important to set
	r.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return m, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return m, err
	}
	return parseJSON[T](body)
}

func toJSON(T any) ([]byte, error) {
	return json.Marshal(T)
}
