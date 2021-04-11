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

//IOC
func TestFunctionAsParameter(t *testing.T) {
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

func TestFunctionReturnFunction2(t *testing.T) {
	result := mathExpression("mult")(2, 2)
	if result != 4 {
		t.Error("Error", result)
	}

	result = mathExpression("add")(2, 2)
	if result != 4 {
		t.Error("Error", result)
	}
}

func TestFunctionReturnFunctionStateful(t *testing.T) {
	//incrementer has a variable, which is can be updated by the returned function
	var doIncrement = incrementer()
	result := doIncrement()
	if result != 2 {
		t.Error("test error", result)
	}

	result = doIncrement()
	if result != 3 {
		t.Error("test error", result)
	}

	result = doIncrement()
	if result != 4 {
		t.Error("test error", result)
	}
}

func TestFunctionBadStateInAnonymous(t *testing.T) {
	var funcs []func() int
	for i := 0; i < 10; i++ {
		funcs = append(funcs, func() int {
			x := i + 1
			return x
		})
	}

	//be careful!!, i is update every time, and the last value is 11 (10 + 1)
	for _, f := range funcs {
		result := f()
		if result != 11 {
			t.Error("test error", result)
		}
	}

	funcs = nil
	for i := 0; i < 10; i++ {
		copyI := i
		funcs = append(funcs, func() int {
			x := copyI + 1
			return x
		})
	}

	for i, f := range funcs {
		result := f()
		if result != i+1 {
			t.Error("test error", result)
		}
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

func mathExpression(exp string) func(int, int) int {
	switch exp {
	case "add":
		return add
	case "mult":
		return mult
	default:
		return func(i1, i2 int) int { return 0 }
	}
}

func incrementer() func() int {
	x := 1
	return func() int {
		x++
		return x
	}
}

func add(x, y int) int {
	return x + y
}

func mult(x, y int) int {
	return x * y
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
