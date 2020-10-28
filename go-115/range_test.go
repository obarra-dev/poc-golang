package main_test

import (
	"testing"
)

func TestRangeWithIndexValue(t *testing.T) {
	arr := [3]int{1, 2, 3}
	var total int
	for i, element := range arr {
		total = total + i + element
	}
	if total == 9 {
		t.Log("OK")
	} else {
		t.Error("Error", total)
	}
}

func TestRangeWithValue(t *testing.T) {
	arr := [3]int{1, 2, 3}
	var total int

	for _, element := range arr {
		total = total + element
	}

	if total == 6 {
		t.Log("OK")
	} else {
		t.Error("Error", total)
	}
}

func TestRangeWithIndex(t *testing.T) {
	arr := [3]int{1, 2, 3}
	var total int

	for i, _ := range arr {
		total = total + i
	}

	if total == 3 {
		t.Log("OK")
	} else {
		t.Error("Error", total)
	}
}

func TestRangeWithOnlyIndex(t *testing.T) {
	arr := [3]int{1, 2, 3}
	var total int

	for i := range arr {
		total = total + i
	}

	if total == 3 {
		t.Log("OK")
	} else {
		t.Error("Error", total)
	}
}
