package main_test

import (
	"fmt"
	"testing"
)

func TestIfWithShortStatement(t *testing.T) {
	if der := "golang"; der == "golang" {
		t.Log("OK")
	} else {
		t.Error("Error")
		t.Fail()
	}

	if der := "golang"; der == "java" {
		t.Error("Error")
		t.Fail()
	} else {
		fmt.Print(der)
		t.Log("OK")
	}
}

func TestIf(t *testing.T) {
	var message string
	var herAge = 29
	if herAge >= 29 {
		message = "Major than 29"
	} else if herAge == 29 {
		message = "Equal 29"
	} else {
		message = "Major than menor than 14"
	}

	if message == "Equal 29" {
		t.Log("OK")
	} else {
		t.Error("Error")
		t.Fail()
	}

}
