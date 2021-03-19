package main_test

import (
	"testing"
)

func TestPointerTo(t *testing.T) {
	name := "omar"
	var pointerToName *string
	pointerToName = &name
	otherPointerToName := &name

	if *pointerToName != "omar" || *otherPointerToName != "omar" {
		t.Error("Error")
	}
}

func TestPointerEqual(t *testing.T) {
	name := "omar"
	var pointerToName *string
	pointerToName = &name

	if pointerToName != &name {
		t.Error("Error")
	}
}

func TestPointerZeroValue(t *testing.T) {
	var pointerToName *string
	if pointerToName != nil {
		t.Error("Error")
	}
}

//new can be used to allocate values
func TestPointerNewWithType(t *testing.T) {
	pointerToName := new(string)

	//pointerToName != "" compilation error
	if *pointerToName != ""  {
		t.Error("Error", pointerToName, *pointerToName)
	}
}

