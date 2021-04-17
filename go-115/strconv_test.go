package main

import (
	"fmt"
	"strconv"
	"testing"
)

// you can not concat string with number, der = "test" + 40 is compile error
func TestStringSPrintf(t *testing.T) {
	result := fmt.Sprintf("%s%+8d", "test", 40)
	if result != "test     +40" {
		t.Error("Error:", result)
	}
}

//Atoi is equivalent to ParseInt(s, 10, 0)
func TestStringStrconv(t *testing.T) {
	number, _ := strconv.Atoi("998")

	if number != 998 {
		t.Error("Error", number)
	}
}

//Use strconv.ParseInt to parse a decimal string (base 10) and check if it fits into an int64.
func TestStringParseInt(t *testing.T) {
	number, err := strconv.ParseInt("-128", 10, 8)
	if number != -128 || err != nil {
		t.Error("Error", number, err)
	}

	number, err = strconv.ParseInt("-129", 10, 8)
	if number != -128 || err.Error() != `strconv.ParseInt: parsing "-129": value out of range` {
		t.Error("Error", number, err.Error())
	}
}

func TestIntToString(t *testing.T) {
	result := strconv.Itoa(404)
	if result != "404" {
		t.Error("Error", result)
	}
}
