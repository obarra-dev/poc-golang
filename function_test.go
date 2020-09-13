package main_test

import "testing"

func TestChangeString(t *testing.T) {
	completeName := "Omar"

	changeString(&completeName)

	if completeName == "Alberto" {
		t.Log("test ok")
	} else {
		t.Error("test Error")
		t.Fail()
	}
}

func changeString(reference *string) {
	*reference = "Alberto"
}
