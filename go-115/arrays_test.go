package main_test

import (
	"testing"
)

//arrays cannot be resized.
func TestArraysDeclareVariable(t *testing.T) {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	var total int
	for d := 0; d < len(arr); d++ {
		total += arr[d]
	}

	if total == 6 {
		t.Log("OK")
	} else {
		t.Error("Error")
		t.Fail()
	}
}

func TestArraysInit(t *testing.T) {
	arr := [3]int{1, 2, 3}
	var total int
	for d := 0; d < len(arr); d++ {
		total += arr[d]
	}

	if total == 6 {
		t.Log("OK")
	} else {
		t.Error("Error")
		t.Fail()
	}
}

func TestArrays2D(t *testing.T) {
	arr2D := [3][2]int{{1, 2}, {4, 3}, {5, 6}}
	var total int
	for d := 0; d < len(arr2D); d++ {
		total += arr2D[d][0]
	}

	if total == 10 {
		t.Log("OK")
	} else {
		t.Error("Error", total)
		t.Fail()
	}
}
