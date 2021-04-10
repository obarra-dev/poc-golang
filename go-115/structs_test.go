package main

import (
	"testing"
)

//The DeeplyEqual() method is defined under “reflect” package.
func TestStructEqualForSimpleType(t *testing.T) {
	point := Point{y: 4, x: 9}
	otherPoint := Point{y: 4, x: 9}
	if point != otherPoint {
		t.Error("Error")
	}
}

func TestStructNotEqualForComplexType(t *testing.T) {
	circle := Circle{&Point{y: 4, x: 9}, 3.3}
	otherCircle := Circle{&Point{y: 4, x: 9}, 3.3}
	if circle == otherCircle {
		t.Error("Error")
	}
}

func TestStructPointerAndCopy(t *testing.T) {
	point := Point{y: 4, x: 9}
	//create a copy
	copy := point
	copy.x = 10

	//create a pointer to point
	pointPointer := &point
	pointPointer.y = 10
	if copy.x == 10 && point.x == 9 && point.y == 10 {
		t.Log("OK")
	} else {
		t.Error("Error")
	}
}

func TestStructLiteral(t *testing.T) {
	pointNamed := Point{y: 4, x: 9}
	point := Point{4, 9}
	pointOne := Point{y: 1}
	zeroPoint := Point{}

	if pointNamed.x == 9 && pointOne.y == 1 && point.x == 4 && zeroPoint.x == 0 {
		t.Log("OK")
	} else {
		t.Error("Error")
	}
}

func TestStructSimple(t *testing.T) {
	circle := Circle{&Point{y: 4, x: 9}, 3.3}
	circle.center.x = 10
	if circle.center.x == 10 && circle.center.y == 4 &&
		circle.redius == 3.3 {
		t.Log("OK")
	} else {
		t.Error("Error")
	}
}

func TestStructFlatten(t *testing.T) {
	circle := CircleFlatten{&Point{4, 9}, 3.3}
	circle.x = 10
	if circle.x == 10 && circle.y == 9 &&
		circle.redius == 3.3 {
		t.Log("OK")
	} else {
		t.Error("Error", circle.x, circle.y)
	}
}
