package main

import (
	"reflect"
	"testing"
)

type animal struct {
	name string
}

func TestReflectTypeOf(t *testing.T) {
	animal := animal{name: "rocki"}
	var tp reflect.Type = reflect.TypeOf(animal)
	if tp.String() != "main.animal" {
		t.Error("Error ", tp.String())
	}

	if tp.Name() != "animal" {
		t.Error("Error ", tp.Name())
	}
}

func TestReflectValue(t *testing.T) {
	animal := animal{name: "rocki"}
	var v reflect.Value = reflect.ValueOf(animal)
	if v.Kind() != reflect.Struct {
		t.Error("Error ", v.Kind())
	}
}

func TestReflectField(t *testing.T) {
	animal := animal{name: "rocki"}
	var v reflect.Value = reflect.ValueOf(animal).Field(0)
	if v.String() != "rocki" {
		t.Error("Error ", v.String())
	}
}
