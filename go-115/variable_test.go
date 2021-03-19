package main_test

import (
	"fmt"
	"math"
	"testing"
)

func TestDeclareAListOfVariableShareType(t *testing.T) {
	//explicit declaration = var declaration
	var c, python bool
	c = true
	python = false
	if !c || python != false {
		t.Error("test Error")
	}
}

// global scope
var java, golang int = 1, 2

func TestVariablesWithInitializers(t *testing.T) {
	var valJava, valGolang = true, "SI!"

	if !(java == 1 && golang == 2 && valJava == true && valGolang == "SI!") {
		t.Error("test Error")
	}
}

func TestShortVariableDeclaration(t *testing.T) {
	//implicit declaration = short declartion, type inference
	valJava, valGolang := true, "SI!"
	kotlin := "yes"

	if kotlin != "yes" || valJava != true || valGolang != "SI!" {
		t.Error("test Error")
	}
}

func TestTypesAndZeroValue(t *testing.T) {
	var name string
	var age int64
	var gender uint8
	var price float64
	var noProgrammer bool
	fmt.Printf("%v %v %v %v %q\n", age, gender, price, noProgrammer, name)
	fmt.Println("this is a test")
	if !(name == "" && age == 0 && gender == 0 && price == 0 && noProgrammer == false) {
		t.Error("test Error")
	}
}

func TestDeclareIntoBlock(t *testing.T) {
	var (
		ToBe   bool   = true
		MaxInt uint64 = 1<<64 - 1
	)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	if MaxInt != 18446744073709551615 || ToBe != true {
		t.Error("test Error")
	}
}

func TestTypeConversions(t *testing.T) {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z, f)
	if x == 3 && y == 4 && f == 5 && z == 5 {
		t.Log("test ok")
	} else {
		t.Error("test Error")
	}
}
