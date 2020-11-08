package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringSplit(t *testing.T) {
	myString := "/api/products/4/cbus/products/other"

	segments := strings.Split(myString, "/products")

	if len(segments) == 3 {
		t.Log("OK")
	} else {
		t.Error("Error", segments, len(segments))
	}
}

func TestStringStrconv(t *testing.T) {
	number, _ := strconv.Atoi("998")

	if number == 998 {
		t.Log("OK")
	} else {
		t.Error("Error", number)
	}
}
