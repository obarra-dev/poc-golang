package main_test

import (
	"testing"
	"fmt"
)

// The zero value of an uninitialized array is pre-filled with the zero value of the arrays declared elemen
func TestArrayZeroValue(t *testing.T) {
	var zeroArray [2]string
	fmt.Println(zeroArray, len(zeroArray), cap(zeroArray))

	if zeroArray != [2]string{"", ""} || len(zeroArray) != 2 || cap(zeroArray) != 2 {
		t.Error("Error", len(zeroArray), cap(zeroArray))
	}
}

// Arrays are values, assigning one array to another copies all the elements
// If you pass an array to a function, it will receive a copy of the array, not a pointer to it
// The size of an array is part of its type. The types [10] int and [20] int are distinct
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

func TestStaticArrays2D(t *testing.T) {
	arr2D := [3][2]int{{1, 2}, {4, 3}, {5, 6}}
	var total int
	for d := 0; d < len(arr2D); d++ {
		total += arr2D[d][0]
	}

	if total != 10 {
		t.Error("Error", total)
	}
}

func TestStaticArrays2DFill(t *testing.T) {
	var arr2D [3][3]int
	fillMatrixWith1(&arr2D)
	var total int
	for d := 0; d < len(arr2D); d++ {
		total += arr2D[d][0]
	}

	if total != 3 {
		t.Error("Error", total)
	}
}

func TestStaticArrays2DFillUsingRange(t *testing.T) {
	var arr2D [3][3]int
	fillMatrixUsingRangeWith1(&arr2D)
	var total int
	for d := 0; d < len(arr2D); d++ {
		total += arr2D[d][0]
	}

	if total != 3 {
		t.Error("Error", total)
	}
}

func fillMatrixUsingRangeWith1(arr2D *[3][3]int) {
	for i, arr := range arr2D {
		for j := range arr {
			arr2D[i][j] = 1
		}
	}
}

func fillMatrixWith1(arr2D *[3][3]int) {
	// number of rows in s2
	n := len(arr2D)

	// number of columns in s2
	m := len(arr2D[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			arr2D[i][j] = 1
		}
	}
}
