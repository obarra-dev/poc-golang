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

	if total != 3 || slice[0] != arr[2] || arr[2] != 6 {
		t.Error("Error", len(slice), slice)
	}
}

//A Slice has both a length and a capacity
// The length of a slice is the number of elements it contains.
// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
func TestSlicesCapAndLen(t *testing.T) {
	arr := [4]int{8, 6, 3, 4}
	slice := arr[2:3]

	if len(slice) != 1 || cap(slice) != 2 {
		t.Error("Error", len(slice), cap(slice))
	}
}

//A slice literal is like an array literal without the length.
func TestSlicesCapAndLenWhenAreEqual(t *testing.T) {
	slice := []int{8, 6, 3, 4}

	if len(slice) != 4 || cap(slice) != 4 {
		t.Error("Error", len(slice), cap(slice))
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

// The zero value of an uninitialized slice is the reference value of nil
func TestSliceZeroValue(t *testing.T) {
	var zeroSlice []string
	fmt.Println(zeroSlice, len(zeroSlice), cap(zeroSlice))

	if zeroSlice != nil || len(zeroSlice) != 0 || cap(zeroSlice) != 0 {
		t.Error("Error", len(zeroSlice), cap(zeroSlice))
	}
}

