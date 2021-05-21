package main_test

import (
	"fmt"
)

func addString(a, b int) string {
	return fmt.Sprintf("%d + %d = %d", a, b, a+b)
}

func Example_addString() {
	word := addString(3, 5)
	fmt.Println(word)
	// Output: 3 + 5 = 8
}

func Example_addString_second() {
	word := addString(5, 5)
	fmt.Println(word)
	// Output: 5 + 5 = 10
}

func Example_addString_third() {
	fmt.Println(addString(5, 5))
	fmt.Println(addString(5, 6))
	fmt.Println(addString(5, 7))
	fmt.Println(addString(5, 8))
	// Output: 5 + 5 = 10
	// 5 + 6 = 11
	// 5 + 7 = 12
	// 5 + 8 = 13
}
