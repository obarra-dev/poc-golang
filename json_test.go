package main

import (
	"encoding/json"
	"testing"
)

func TestJsonMarshal(t *testing.T) {
	m := Message{"Maru", "Hello", 1294706395881547000}
	b, _ := json.Marshal(m)

	if string(b) == `{"Name":"Maru","Body":"Hello","Time":1294706395881547000}` {
		t.Log("OK")
	} else {
		t.Error("Error", string(b))
		t.Fail()
	}
}

func TestJsonUnMarshal(t *testing.T) {
	b := []byte(`{"Name":"Maru","Body":"Hello","Time":1294706395881547000}`)

	var m Message
	json.Unmarshal(b, &m)

	if m.Name == "Maru" {
		t.Log("OK")
	} else {
		t.Error("Error")
		t.Fail()
	}
}

type Message struct {
	Name string
	Body string
	Time int64
}
