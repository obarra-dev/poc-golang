package main

import (
	"testing"
)

func TestMethodInCustomTypeSimple(t *testing.T) {
	language := Language("java")

	if name := language.GetName(); name == "java" {
		t.Log("OK")
	} else {
		t.Error("Error", name)
		t.Fail()
	}
}

func TestMethodAbs(t *testing.T) {
	point := Point{4, 8}
	pow := point.Abs()
	if pow == 8.94427190999916 {
		t.Log("OK")
	} else {
		t.Error("Error", pow)
		t.Fail()
	}
}

//pointer receiver
// the method can modify the value that its receiver points to
// avoid copying the value on each method call
func TestMethodScale(t *testing.T) {
	point := Point{4, 8}
	point.Scale(2)
	if point.x == 8 && point.y == 16 {
		t.Log("OK")
	} else {
		t.Error("Error", point)
		t.Fail()
	}
}

func TestMethodIndirectionValues(t *testing.T) {
	point := &Point{4, 8}
	//is interpreted as (*point).Scale()
	point.Scale(2)
	if point.x == 8 && point.y == 16 {
		t.Log("OK")
	} else {
		t.Error("Error", point)
		t.Fail()
	}
}
