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

func TestStringBuilder(t *testing.T) {
	var builder strings.Builder
	builder.WriteString("First")
	builder.WriteString("Second")
	result := builder.String()

	if result == "FirstSecond" {
		t.Log("OK")
	} else {
		t.Error("Error", result)
	}
}

func TestIntToString(t *testing.T) {
	result := strconv.Itoa(404)
	if result != "404" {
		t.Error("Error", result)
	}
}

func TestStringsJoin(t *testing.T) {
	array := [3]string{"a", "b", "c"}
	result := strings.Join(array[:], ",")

	if result != "a,b,c" {
		t.Error("Error", result)
	}
}
