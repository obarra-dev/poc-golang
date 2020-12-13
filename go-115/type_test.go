package main_test

import (
	"strconv"
	"testing"
)

//custom type is known as deck
type app []string

func TestCustomTypeDeclaration(t *testing.T) {
	apps := app{"Facebook", "Instagram", "WhatsApp"}
	if len(apps) != 3 {
		t.Error("test Error", apps)
	}
}

func TestCustomTypeDeclarationUseForMap(t *testing.T) {
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
