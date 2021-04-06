package main

import (
	"encoding/json"
	"testing"
	"time"
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

//por defecto toma la key es igual a nombre del atributo, y esta debe comenzar en mayuscula.
//para cambiar usar json:"atribute_name"
func TestJsonMarshalString(t *testing.T) {
	p := Person{"omar", 29}
	pJSON, _ := json.Marshal(p)

	if string(pJSON) != `{"name":"omar"}` {
		t.Error("Error", string(pJSON))
	}
}

func TestJsonUnMarshal(t *testing.T) {
	b := []byte(`{"Name":"Maru","Body":"Hello","Time":1294706395881547000}`)

	var m Message
	json.Unmarshal(b, &m)

	if m.Name != "Maru" {
		t.Error("Error")
	}
}

type UserOmmitAlways struct {
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

//if the field tag is "-", the field is always omitted.
func TestJsonUnMarshalWithAlwaysOmited(t *testing.T) {
	b := []byte(`{"name":"Maru","password":"dddd", "createdAt": "2019-09-23T16:08:21.124481-04:00"}`)

	var uOne UserOmmitAlways
	json.Unmarshal(b, &uOne)

	if uOne.Name != "Maru" || uOne.Password != "" || uOne.CreatedAt.IsZero() {
		t.Error("Error ", uOne)
	}

	uTwo := UserOmmitAlways{
		Name:     "Maru",
		Password: "dddd",
		CreatedAt: time.Date(
			2019, 04, 04, 20, 34, 58, 651387237, time.UTC),
	}

	uJSON, _ := json.Marshal(uTwo)

	if string(uJSON) != `{"name":"Maru","createdAt":"2019-04-04T20:34:58.651387237Z"}` {
		t.Error("Error ", string(uJSON))
	}
}

type UserOmmitEmpty struct {
	Name      string    `json:"name"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

//omitempty augmentation causes the JSON encoder to skip that field,
func TestJsonUnMarshalWithEmptyOmited(t *testing.T) {
	b := []byte(`{"name":"Maru","password":"dddd", "createdAt": "2019-09-23T16:08:21.124481-04:00"}`)

	var uOne UserOmmitEmpty
	json.Unmarshal(b, &uOne)

	if uOne.Name != "Maru" || uOne.Password != "dddd" || uOne.CreatedAt.IsZero() {
		t.Error("Error ", uOne)
	}

	uTwo := UserOmmitEmpty{
		Name:     "Maru",
		Password: "",
		CreatedAt: time.Date(
			2019, 04, 04, 20, 34, 58, 651387237, time.UTC),
	}

	uJSON, _ := json.Marshal(uTwo)
	if string(uJSON) != `{"name":"Maru","createdAt":"2019-04-04T20:34:58.651387237Z"}` {
		t.Error("Error ", string(uJSON))
	}

	uTwo = UserOmmitEmpty{
		Name:     "Maru",
		Password: "ddd",
		CreatedAt: time.Date(
			2019, 04, 04, 20, 34, 58, 651387237, time.UTC),
	}

	uJSON, _ = json.Marshal(uTwo)
	if string(uJSON) != `{"name":"Maru","password":"ddd","createdAt":"2019-04-04T20:34:58.651387237Z"}` {
		t.Error("Error ", string(uJSON))
	}
}

type Person struct {
	Name string `json:"name"`
	//age is ignored by Marshalling because it is private
	age int `json:"age"`
}
