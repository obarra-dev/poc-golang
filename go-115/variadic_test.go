package main

import (
	"testing"
)

func TestVariadic(t *testing.T) {
	numbers := []int{1, 2, 3, 5, 8, 13}
	result := sum(numbers...)

	if result == 32 {
		t.Log("OK")
	} else {
		t.Error("Error", result)
	}
}

//specifies a length equal to the number of elements in the literal.
func TestVariadicInArrayLiteral(t *testing.T) {
	numbers := [...]int{1, 2, 3, 4, 5, 6}

	if len(numbers) == 6 {
		t.Log("OK")
	} else {
		t.Error("Error", len(numbers))
	}
}

func sum(nums ...int) int {
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}
