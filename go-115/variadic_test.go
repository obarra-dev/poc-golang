package main

import (
	"testing"
)

func TestVariadicPassingSlice(t *testing.T) {
	numbers := []int{1, 2, 3, 5, 8, 13}
	result := sum(numbers...)

	if result != 32 {
		t.Error("Error", result)
	}
}

func TestVariadicPassingParameters(t *testing.T) {
	result := sum(1, 2, 3, 5, 8, 13)
	if result != 32 {
		t.Error("Error", result)
	}
}

// other use of variadic
//specifies a length equal to the number of elements in the literal.
func TestVariadicInArrayLiteral(t *testing.T) {
	numbers := [...]int{1, 2, 3, 4, 5, 6}

	// you can not pass array with specifies lenth into variadic function
	//sum(numbers...)
	if len(numbers) != 6 {
		t.Error("Error", len(numbers))
	}
}

// variadic function
// it has to be the last parameter in the function
func sum(nums ...int) int {
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}
