package main

import (
	"fmt"
	"testing"
)

func TestInterfaceValueAndType(t *testing.T) {
	circle := &Circle{&Point{y: 4, x: 9}, 3.3}
	value, myType := describe(circle)

	if value == "&{0xc0000122d0 3.3}" && myType == "*main.Circle" {
		t.Log("test ok")
	} else {
		t.Error("test Error")
		t.Fail()
	}
}

func TestInterfaceValueWithNil(t *testing.T) {
	var circle *Circle
	var shape shape
	shape = circle

	//if method do not treat a nil it return panic value method xxx called using nil
	flag := shape.hasImplementation()
	if flag == false {
		t.Log("test ok")
	} else {
		t.Error("test Error")
		t.Fail()
	}
}

func TestNilInterfaceValue(t *testing.T) {
	var shape shape
	value, myType := describe(shape)

	if value == "<nil>" && myType == "<nil>" {
		t.Log("test ok")
	} else {
		t.Error("test Error")
		t.Fail()
	}
}

func TestEmptyInterfaceString(t *testing.T) {
	value, myType := describeemptyInterface("hello")

	if value == "hello" && myType == "string" {
		t.Log("test ok")
	} else {
		t.Error("test Error", value, myType)
		t.Fail()
	}
}

func TestEmptyInterfaceInt(t *testing.T) {
	value, myType := describeemptyInterface(404)

	if value == "404" && myType == "int" {
		t.Log("test ok")
	} else {
		t.Error("test Error", value, myType)
		t.Fail()
	}
}

func TestInterfaceTypeAssertion(t *testing.T) {
	var i interface{} = "barra"

	value, ok := i.(string)

	if ok == true && value == "barra" {
		t.Log("test ok")
	} else {
		t.Error("test Error", value, ok)
		t.Fail()
	}
}

func TestInterfaceTypeAssertionSwitch(t *testing.T) {
	var i interface{} = "barra"

	typo := ""
	switch i.(type) {
	case int:
		typo = "int"
	case string:
		typo = "string"
	default:
		typo = "no type"
	}

	if typo == "string" {
		t.Log("test ok")
	} else {
		t.Error("test Error", typo)
		t.Fail()
	}
}

func describeemptyInterface(emptyInterface emptyInterface) (string, string) {
	fmt.Printf("(%v, %T)\n", emptyInterface, emptyInterface)
	var value = fmt.Sprintf("%v", emptyInterface)
	fmt.Println(value)
	var myType = fmt.Sprintf("%T", emptyInterface)
	fmt.Println(myType)
	return value, myType
}

func describe(myShape shape) (string, string) {
	fmt.Printf("(%v, %T)\n", myShape, myShape)
	var value = fmt.Sprintf("%v", myShape)
	fmt.Println(value)
	var myType = fmt.Sprintf("%T", myShape)
	fmt.Println(myType)
	return value, myType
}
