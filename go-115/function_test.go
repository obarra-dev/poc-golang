package main_test

import (
	"fmt"
	"testing"
)

// * dereference pointer operator
func TestPointerAndDereferenceOperator(t *testing.T) {
	completeName := "Omar"

	changeString(&completeName)

	if completeName != "Alberto" {
		t.Error("test Error")
	}
}

func changeString(reference *string) {
	*reference = "Alberto"
}

// avoid to use naming retun values, it is hard to read
func TestFunctionParametersShareATypeAndReturnThreeValuesAndTheReturnValuesAreNamed(t *testing.T) {
	added, subs, mul := performNumbers(4, 2)

	if added != 6 || subs != 2 || mul != 16 {
		t.Error("test error")
	}
}

func TestFunctionAnonimus(t *testing.T) {
	deny := func(x int) int {
		return x * -1
	}

	var nagtiveValue = deny(10)

	if nagtiveValue != -10 {
		t.Error("test error")
	}
}

func TestFunctionIOC(t *testing.T) {
	deny := func(x int) int {
		if x > 0 {
			return x * -1
		}
		return x
	}
	var negativeValue = functionWithFunctionArgument(deny)

	adder := func(x int) int {
		if x < 0 {
			return x * -1
		}
		return x
	}
	var positiveValue = functionWithFunctionArgument(adder)

	if positiveValue != 4 || negativeValue != -4 {
		t.Error("test error")
	}
}

//Function closures
func TestFunctionReturnFunction(t *testing.T) {
	var concater = concat("Maru")
	var result = concater(" :)")
	if result != "Hi Maru :)" {
		t.Error("test error")
	}
}

func TestFunctionType(t *testing.T) {
	result := squareSum(1)(1)(1)
	if result != 3 {
		t.Error("Error", result)
	}
}

type firstFunctionType func(int) int
type secondFunctionType func(int) firstFunctionType

func squareSum(x int) secondFunctionType {
	return func(y int) firstFunctionType {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

//Named Return Values
func performNumbers(x, y int) (w, z, d int) {
	defer fmt.Println("it is finally")
	w = x + y
	z = x - y
	d = x * x
	return
}

func functionWithFunctionArgument(myfunc func(int) int) int {
	return myfunc(4)
}

func concat(word string) func(string) string {
	var newWord = "Hi " + word
	return func(icond string) string {
		return newWord + icond
	}
}
