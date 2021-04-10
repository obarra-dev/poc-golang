package main

import (
	"testing"
)

func TestMethodInCustomTypeSimple(t *testing.T) {
	//conversion like string("java")
	language := Language("java")

	if name := language.GetName(); name != "name: java" {
		t.Error("Error", name)
	}
}

func TestMethodValueBasedReceiver(t *testing.T) {
	point := Point{4, 8}
	pow := point.Abs()
	if pow != 8.94427190999916 {
		t.Error("Error", pow)
	}
}

//pointer receiver
// the method can modify the value that its receiver points to
// avoid copying the value on each method call
func TestMethodPointerBasedReceiver(t *testing.T) {
	point := Point{4, 8}
	point.Scale(2)
	if point.x != 8 || point.y != 16 {
		t.Error("Error", point)
	}
}

func TestMethodIndirectionValues(t *testing.T) {
	point := &Point{4, 8}
	//is interpreted as (*point).Scale()
	point.Scale(2)
	if point.x != 8 || point.y != 16 {
		t.Error("Error", point)
	}
}

func TestMethodAndInterfacesAndReceiver(t *testing.T) {
	//var point ShapeInterface = Point{4, 8}  is a compile error  becaouse Point with pointer receiver implements ShapeInterface
	var point ShapeInterface = &Point{4, 8}
	point.Scale(2)

	if point.Abs() != 17.88854381999832 {
		t.Error("Error", point.Abs())
	}
}
