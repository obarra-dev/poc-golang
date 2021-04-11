package main_test

import (
	"strconv"
	"testing"
)

//Type aliases
type day = int

//day and int are the same
func TestTypeAlias(t *testing.T) {
	var n int = 0
	var a day = n
	if a != 10 {
		t.Error("test Error", a)
	}
}

//Type definition
type month int

//The new type is called a defined type. It is different from any other type, including the type it is created from.
func TestTypeDefinition(t *testing.T) {
	var n int = 10
	//var a month = n  compile error
	var a month = month(n)
	if a != 10 {
		t.Error("test Error", a)
	}
}

//Type definition
type app []string

func TestTypeDefinition2(t *testing.T) {
	apps := app{"Facebook", "Instagram", "WhatsApp"}
	if len(apps) != 3 {
		t.Error("test Error", apps)
	}
}

type Vehicle int

const (
	Bike Vehicle = iota
	Bus
	Train
)

func TestTypeDefinitionConstantsUsingiota(t *testing.T) {
	if Bike != 0 || Bus != 1 || Train != 2 {
		t.Error("test Error", Bike, Bus, Train)
	}
}

func TestTypeDefinitionForMap(t *testing.T) {
	integers := []int{1, 2, 3}
	strings := Map(func(x int) string { return strconv.Itoa(x) }, integers)
	if len(strings) != 3 || strings[0] != "1" {
		t.Error("test Error", strings)
	}
}

type operatorIntString func(x int) string

func Map(op operatorIntString, a []int) []string {
	res := make([]string, len(a))
	for i, x := range a {
		res[i] = op(x)
	}
	return res
}
