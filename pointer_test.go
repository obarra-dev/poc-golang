package main_test

import (
	"testing"
)

func TestPointerTo(t *testing.T) {
	name := "omar"
	var pointerToName *string
	pointerToName = &name
	otherPointerToName := &name

	if *pointerToName == "omar" && *otherPointerToName == "omar" {
		t.Log("OK")
	} else {
		t.Error("Error")
		t.Fail()
	}
}

func TestPointerEqual(t *testing.T) {
	name := "omar"
	var pointerToName *string
	pointerToName = &name

	if pointerToName == &name {
		t.Log("OK")
	} else {
		t.Error("Error")
		t.Fail()
	}
}

func TestPointerZeroValye(t *testing.T) {
	var pointerToName *string

	if pointerToName == nil {
		t.Log("OK")
	} else {
		t.Error("Error")
		t.Fail()
	}
}

func TestPointerNewWithType(t *testing.T) {
	pointerToName := new(string)

	if *pointerToName == "" {
		t.Log("OK")
	} else {
		t.Error("Error")
		t.Fail()
	}
}
