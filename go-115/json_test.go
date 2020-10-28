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
	}
}

//por defecto toma la key es igual a nombre del atributo, y esta debe estar comenzar en mayuscula.
//para cambiar usar json:"atribute_name"
func TestJsonMarshalString(t *testing.T) {
	m := myHandler{"omar"}
	b, _ := json.Marshal(m)

	if string(b) == `{"message":"omar"}` {
		t.Log("OK")
	} else {
		t.Error("Error", string(b))
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
	}
}

type myHandler struct {
	Message string `json:"message"`
}
