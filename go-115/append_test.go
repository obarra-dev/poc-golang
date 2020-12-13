package main_test

import (
	"testing"
)

func TestAppendInANilArray(t *testing.T) {
	var arraWithOutLimit []int
	arraWithOutLimit = append(arraWithOutLimit, 10, 12, 13)
	if len(arraWithOutLimit) != 3 || cap(arraWithOutLimit) != 4 {
		t.Error("Error", len(arraWithOutLimit), cap(arraWithOutLimit))
	}
}

func TestAppendInArrayWithValues(t *testing.T) {
	arraWithOutLimit := []int{3, 4}
	arraWithOutLimit = append(arraWithOutLimit, 10, 12, 13)
	if len(arraWithOutLimit) != 5 || cap(arraWithOutLimit) != 6 {
		t.Error("Error", len(arraWithOutLimit), cap(arraWithOutLimit))
	}
}

//Capacity  depends on the size of the elements stored in the array
func TestCapacityWhenIsString(t *testing.T) {
	arraWithOutLimit := []string{}
	if cap(arraWithOutLimit) != 0 {
		t.Error("Error", cap(arraWithOutLimit))
	}
	arraWithOutLimit = append(arraWithOutLimit, "1")
	if cap(arraWithOutLimit) != 1 {
		t.Error("Error", cap(arraWithOutLimit))
	}
	arraWithOutLimit = append(arraWithOutLimit, "2")
	if cap(arraWithOutLimit) != 2 {
		t.Error("Error", cap(arraWithOutLimit))
	}
	arraWithOutLimit = append(arraWithOutLimit, "3")
	if cap(arraWithOutLimit) != 4 {
		t.Error("Error", cap(arraWithOutLimit))
	}
}

func TestCapacityWhenIsSomeStruct(t *testing.T) {
	type str3Bytes struct {
		a byte
		b byte
		c byte
	}

	arraWithOutLimit := []str3Bytes{}
	if cap(arraWithOutLimit) != 0 {
		t.Error("Error", cap(arraWithOutLimit))
	}
	arraWithOutLimit = append(arraWithOutLimit, str3Bytes{1, 2, 3})
	if cap(arraWithOutLimit) != 2 {
		t.Error("Error", cap(arraWithOutLimit))
	}
	arraWithOutLimit = append(arraWithOutLimit, str3Bytes{1, 2, 3})
	if cap(arraWithOutLimit) != 2 {
		t.Error("Error", cap(arraWithOutLimit))
	}
	arraWithOutLimit = append(arraWithOutLimit, str3Bytes{1, 2, 3})
	if cap(arraWithOutLimit) != 5 {
		t.Error("Error", cap(arraWithOutLimit))
	}
}
