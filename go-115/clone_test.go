package main

import (
	"testing"
)

func TestCloneSimple(t *testing.T) {
	point := Point{y: 4, x: 9}
	otherPoint := point
	if point == otherPoint {
		t.Log("OK")
	} else {
		t.Error("Error")
	}
}

func TestCloneSimpleWhenChangeXvalue(t *testing.T) {
	point := Point{y: 4, x: 9}
	otherPoint := point
	otherPoint.x = 10
	if point != otherPoint && point.x == 9 && otherPoint.x == 10 {
		t.Log("OK")
	} else {
		t.Error("Error")
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
