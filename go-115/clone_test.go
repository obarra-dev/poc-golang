package main

import (
	"testing"
)

func TestCopySlice(t *testing.T) {
	x := [3]string{"zero", "one", "two"}

	y := make([]string, len(x))
	copy(y, x[:]) // slice "z" is a copy of the slice created from array "x"

	y[1] = "none" // the value at index 1 is now "Belka" for only "y"

	if x[1] != "one" || y[1] != "none" {
		t.Error("Error: ", y)
	}
}

func TestCloneSimpleSlice(t *testing.T) {
	x := [3]string{"zero", "one", "two"}

	y := x[:] // slice "y" points to the underlying array "x"

	y[1] = "none" // the value at index 1 is now "Belka" for both "y" and "x"

	if x[1] != "none" || y[1] != "none" {
		t.Error("Error: ", y)
	}
}

//Go supports equality checking structs.
func TestCloneSimple(t *testing.T) {
	point := Point{y: 4, x: 9}

	// make a copy of the value, so ponter is other
	otherPoint := point
	if point != otherPoint || &point == &otherPoint {
		t.Error("Error", &point, &otherPoint)
	}
}

func TestCloneSimpleWhenChangeXvalue(t *testing.T) {
	point := Point{y: 4, x: 9}
	otherPoint := point
	otherPoint.x = 10
	if point == otherPoint || point.x != 9 || otherPoint.x != 10 {
		t.Error("Error", &point, &otherPoint)
	}
}

func TestCloneComplex(t *testing.T) {
	nick := "barra"
	complex := ComplexStruct{"omar",
		29,
		&nick,
		[]string{"los", "jazmines"},
		map[string]string{"home": "1212132"}}

	otherComplex := complex
	*otherComplex.nick = "test"
	if *complex.nick == "test" && *otherComplex.nick == "test" {
		t.Log("OK")
	} else {
		t.Error("Error", otherComplex)
	}
}

func TestCloneComplexWithCreatCopy(t *testing.T) {
	nick := "barra"
	complex := ComplexStruct{"omar",
		29,
		&nick,
		[]string{"los", "jazmines"},
		map[string]string{"home": "1212132"}}

	otherComplex := complex.createCopy()
	*otherComplex.nick = "test"
	if *complex.nick == "barra" && *otherComplex.nick == "test" {
		t.Log("OK")
	} else {
		t.Error("Error", otherComplex)
	}
}
