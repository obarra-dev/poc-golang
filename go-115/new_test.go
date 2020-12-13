package main

import (
	"reflect"
	"testing"
)

// new can be used to ALLOCATES values
// new return a pointer
func TestNewWhenAllocateString(t *testing.T) {
	pointerString := new(string)

	if *pointerString != "" {
		t.Error("Error", pointerString, *pointerString)
	}

	if reflect.ValueOf(pointerString).Kind() != reflect.Ptr {
		t.Error("Error", pointerString, *pointerString)
	}
}

func TestNewWhenAllocateInt(t *testing.T) {
	pointerInt := new(int)
	// d := &int // Illegal
	if *pointerInt != 0 {
		t.Error("Error", pointerInt, *pointerInt)
	}

	if reflect.ValueOf(pointerInt).Kind() != reflect.Ptr {
		t.Error("Error", pointerInt, *pointerInt)
	}

	// Works, but it is less convenient to write than new(int)
	var integer int

	if reflect.ValueOf(&integer).Kind() != reflect.Ptr {
		t.Error("Error", &integer, *&integer)
	}
}

func TestNewWhenAllocateStruct(t *testing.T) {
	//	&Point{}      // OK = new(Point)
	//&Point{2, 3}  // Combines allocation and initialization
	pointerPoint := new(Point)
	pointerPoint.x = 10
	pointerPoint.y = 11
	if pointerPoint.x != 10 || pointerPoint.y != 11 {
		t.Error("Error", pointerPoint, *pointerPoint)
	}

	if reflect.ValueOf(pointerPoint).Kind() != reflect.Ptr {
		t.Error("Error", pointerPoint, *pointerPoint)
	}
}

func TestNoUseNewWhenAllocateStruct(t *testing.T) {
	pointerPoint := &Point{}
	pointerPoint.x = 10
	pointerPoint.y = 11
	if pointerPoint.x != 10 || pointerPoint.y != 11 {
		t.Error("Error", pointerPoint, *pointerPoint)
	}

	if reflect.ValueOf(pointerPoint).Kind() != reflect.Ptr {
		t.Error("Error", pointerPoint, *pointerPoint)
	}
}

func TestNewVsMakeUsingNew(t *testing.T) {
	//new return a pointer
	arrayPointer := new([]int)

	if reflect.ValueOf(arrayPointer).Kind() != reflect.Ptr {
		t.Error("Error", arrayPointer, *arrayPointer)
	}

	//new no intilize, so the internal value is nil
	if arrayPointer == nil || *arrayPointer != nil {
		t.Error("Error", arrayPointer, *arrayPointer)
	}
}

func TestNewVsMakeUsingMake(t *testing.T) {
	//make return a slice, no a pointer
	array := make([]int, 0)
	arrayPointer := &array

	if reflect.ValueOf(array).Kind() != reflect.Slice {
		t.Error("Error", array, reflect.ValueOf(array).Kind())
	}

	if reflect.ValueOf(arrayPointer).Kind() != reflect.Ptr {
		t.Error("Error", arrayPointer, *arrayPointer)
	}

	//make intilize, so the internal value has value, it is a slice
	if arrayPointer == nil || *arrayPointer == nil {
		t.Error("Error", arrayPointer, *arrayPointer)
	}
}
