package main_test

import (
	"testing"
)

//  parentheses are redundant, so are optional
func TestIf(t *testing.T) {
	var message string
	var herAge = 29
	if herAge > 29 {
		message = "Major than 29"
	} else if herAge == 29 {
		message = "Equal 29"
	} else {
		message = "Major than menor than 14"
	}

	if message != "Equal 29" {
		t.Error("Error", message)
	}
}

//if and switch accept an optional initialization statement
func TestIfWithShortStatement(t *testing.T) {
	if der := "golang"; der != "golang" {
		t.Error("Error")
	}
	if der := "golang"; der != "java" {
		t.Log("OK")
	}
}
