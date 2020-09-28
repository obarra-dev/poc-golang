package main_test

import (
	"fmt"
	"testing"
)

func TestLogicalOperator(t *testing.T) {
	isTrue := (false || true) && !false && 10-5 == 5
	if isTrue == true {
		t.Log("OK")
	} else {
		t.Error("Error")
		t.Fail()
	}
}

func TestArithmeticOperator(t *testing.T) {
	var number float32 = 1
	var otherNumber float32 = 3
	resultFloat := number / otherNumber
	fmt.Printf("%g", resultFloat)

	var numberint int = 1
	var otherNumberInt int = 3
	resultInt := numberint / otherNumberInt
	fmt.Printf("%d", resultInt)

	if resultInt == 0 && resultFloat == 0.33333334 {
		t.Log("OK")
	} else {
		t.Error("Error,", resultFloat, resultInt)
		t.Fail()
	}
}
