package main

import (
	"encoding/json"
	"testing"
)

func TestJsonMarshal(t *testing.T) {
	m := Message{"Maru", "Hello", 1294706395881547000}
	mJSON, _ := json.Marshal(m)

	if string(mJSON) == `{"Name":"Maru","Body":"Hello","Time":1294706395881547000}` {
		t.Log("OK")
	} else {
		t.Error("Error", string(mJSON))
	}
}

//por defecto toma la key es igual a nombre del atributo, y esta debe estar comenzar en mayuscula.
//para cambiar usar json:"atribute_name"
func TestJsonMarshalString(t *testing.T) {

	p := Person{"omar", 29}
	pJSON, _ := json.Marshal(p)

	if string(pJSON) == `{"name":"omar"}` {
		t.Log("OK")
	} else {
		t.Error("Error", string(pJSON))
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

type Person struct {
	Name string `json:"name"`
	//age is ignored by Marshalling
	age int `json:"age"`
}
