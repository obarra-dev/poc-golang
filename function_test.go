package main_test

import (
	"fmt"
	"testing"
)

func TestPointerAndDereferenceOperator(t *testing.T) {
	completeName := "Omar"

	changeString(&completeName)

	if completeName == "Alberto" {
		t.Log("test ok")
	} else {
		t.Error("test Error")
		t.Fail()
	}
}

func changeString(reference *string) {
	*reference = "Alberto"
}

func TestFunctionParametersShareATypeAndReturnThreeValuesAndTheReturnValuesAreNamed(t *testing.T) {
	added, subs, mul := performNumbers(4, 2)

	if added == 6 && subs == 2 && mul == 16 {
		t.Log("test ok")
	} else {
		t.Error("test error")
		t.Fail()
	}
}

func TestFunctionAnonimus(t *testing.T) {
	deny := func(x int) int {
		return x * -1
	}

	var nagtiveValue = deny(10)

	if nagtiveValue == -10 {
		t.Log("test ok")
	} else {
		t.Error("test error")
		t.Fail()
	}
}

func TestFunctionIOC(t *testing.T) {
	deny := func(x int) int {
		if x > 0 {
			return x * -1
		}
		return x
	}
	var negativeValue = ioc(deny)

	adder := func(x int) int {
		if x < 0 {
			return x * -1
		}
		return x
	}
	var positiveValue = ioc(adder)

	if positiveValue == 4 && negativeValue == -4 {
		t.Log("test ok")
	} else {
		t.Error("test error")
		t.Fail()
	}
}

func TestFunctionReturnFunction(t *testing.T) {
	var result = returnFunction("Maru")()
	fmt.Println(result)
	if result == "Hi Maru :)" {
		t.Log("test ok")
	} else {
		t.Error("test error")
		t.Fail()
	}
}

func performNumbers(x, y int) (w, z, d int) {
	defer fmt.Println("it is finally")
	w = x + y
	z = x - y
	d = x * x
	return
}

func ioc(myfunc func(int) int) int {
	return myfunc(4)
}

func returnFunction(word string) func() string {
	var newWord = "Hi " + word
	return func() string {
		return newWord + " :)"
	}
}
