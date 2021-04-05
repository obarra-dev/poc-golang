package main

import (
	"fmt"
	"testing"
)

func TestInterfaceValueAndType(t *testing.T) {
	circle := &Circle{&Point{y: 4, x: 9}, 3.3}
	value, myType := describe(circle)

	//in this case value is a string with pointer value of this interface
	if value != "" && myType == "*main.Circle" {
		t.Log("test ok")
	} else {
		t.Error("test Error", value, myType)
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
	}
}

func TestNilInterfaceValue(t *testing.T) {
	var shape shape
	value, myType := describe(shape)

	if value == "<nil>" && myType == "<nil>" {
		t.Log("test ok")
	} else {
		t.Error("test Error")
	}
}

func TestEmptyInterfaceString(t *testing.T) {
	value, myType := describeEmptyInterface("hello")

	if value == "hello" && myType == "string" {
		t.Log("test ok")
	} else {
		t.Error("test Error", value, myType)
	}
}

func TestEmptyInterfaceInt(t *testing.T) {
	value, myType := describeEmptyInterface(404)

	if value == "404" && myType == "int" {
		t.Log("test ok")
	} else {
		t.Error("test Error", value, myType)
	}
}

func TestInterfaceTypeAssertion(t *testing.T) {
	var i interface{} = "barra"

	value, ok := i.(string)

	if ok == true && value == "barra" {
		t.Log("test ok")
	} else {
		t.Error("test Error", value, ok)
	}
}

//Type switch
func TestInterfaceTypeAssertionSwitch(t *testing.T) {
	var i interface{} = "barra"

	typo := ""
	switch i.(type) {
	case bool:
		typo = "bool"
	case *bool:
		typo = "*bool"
	case *int:
		typo = "*int"
	case int:
		typo = "int"
	case string:
		typo = "string"
	default:
		typo = "no type"
	}

	if typo != "string" {
		t.Error("test Error", typo)
	}
}

func describeEmptyInterface(emptyInterface emptyInterface) (string, string) {
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
