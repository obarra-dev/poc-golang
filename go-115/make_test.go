package main_test

import (
	"fmt"
	"testing"
)

//Make is a build-in function, it is how you create dynamically-sized arrays.
// The make built-in function allocates and initializes an object of type slice, map, or chan (only).
func TestMake(t *testing.T) {
	//The make function allocates a zeroed array and returns a slice that refers to that array:
	maked := make([]int, 5)
	if len(maked) != 5 || cap(maked) != 5 {
		t.Error("Error", len(maked), cap(maked))
	}

	count := 0
	for _, v := range maked {
		if v != 0 {
			t.Error("Error", v)
		}
		count++
	}

	if count != 5 {
		t.Error("Error", count)
	}
}

func TestMakeCapAndLengDiff(t *testing.T) {
	maked := make([]int, 0, 5)
	if len(maked) != 0 || cap(maked) != 5 {
		t.Error("Error", len(maked), cap(maked))
	}

	count := 0
	for _, v := range maked {
		if v != 0 {
			t.Error("Error", v)
		}
		count++
	}

	if count != 0 {
		t.Error("Error", count)
	}
}

func TestMakeMap(t *testing.T) {
	maked := make(map[string]int)
	if len(maked) != 0 {
		t.Error("Error", len(maked))
	}
}

func TestDynamicMatrix(t *testing.T) {
	//matrix := [][]int{}

	//create the rows
	matrix := make([][]int, 4)

	//create the columns
	for i := range matrix {
		matrix[i] = make([]int, 3)
	}

	fillDynamicMatrixWith1(matrix)

	fmt.Println(len(matrix), len(matrix[0]), matrix[0][0])

	showMatrix(matrix)

	var total int
	for d := 0; d < len(matrix); d++ {
		total += matrix[d][0]
	}

	if total != 4 {
		t.Error("Error", total)
	}
}

func showMatrix(matrix [][]int) {
	for _, v := range matrix {
		for _, b := range v {
			fmt.Print(b)
		}
		fmt.Println()
	}
}

func fillDynamicMatrixWith1(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = 1
		}
	}
}
