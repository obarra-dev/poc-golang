package main_test

import (
	"fmt"
	"testing"
)

func TestSlicesAreaReferencesToArrays(t *testing.T) {
	arr := [4]int{8, 6, 3, 4}
	slice := arr[2:3]

	var total int
	for d := 0; d < len(slice); d++ {
		total += slice[d]
	}

	//index 0 of slice is index 2 array
	slice[0] = 6

	if total == 3 && arr[2] == 6 {
		t.Log("OK")
	} else {
		t.Error("Error", len(slice))
		t.Fail()
	}
}

func TestSlicesCapAndLen(t *testing.T) {
	arr := [4]int{8, 6, 3, 4}
	slice := arr[2:3]

	if len(slice) == 1 && cap(slice) == 2 {
		t.Log("OK")
	} else {
		t.Error("Error", len(slice), cap(slice))
		t.Fail()
	}
}

func TestSlicesCapAndLenNew(t *testing.T) {
	arraWithOutLimit := []int{3, 4}
	if len(arraWithOutLimit) == 2 && cap(arraWithOutLimit) == 2 {
		t.Log("OK")
	} else {
		t.Error("Error", len(arraWithOutLimit), cap(arraWithOutLimit))
		t.Fail()
	}
}

//A slice literal is like an array literal without the length.
func TestSlicesLiteral(t *testing.T) {
	slice := []int{8, 6, 3, 4}

	if len(slice) == 4 && cap(slice) == 4 {
		t.Log("OK")
	} else {
		t.Error("Error", len(slice), cap(slice))
		t.Fail()
	}
}

func TestSlicesDefaultValue(t *testing.T) {
	slice := []int{8, 6, 3, 4}

	complete := slice[:]
	fmt.Println(complete)

	until := slice[:2]
	fmt.Println(until)

	since := slice[2:]
	fmt.Println(since)

	otherComplete := slice[0:4]
	fmt.Println(otherComplete)
}

func TestSliceNil(t *testing.T) {
	var nilSlices []string
	fmt.Println(nilSlices, len(nilSlices), cap(nilSlices))

	if nilSlices == nil && len(nilSlices) == 0 && cap(nilSlices) == 0 {
		t.Log("OK")
	} else {
		t.Error("Error", len(nilSlices), cap(nilSlices))
		t.Fail()
	}
}
