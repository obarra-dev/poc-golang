package main

import "testing"

func TestStructEqualZeroValue(t *testing.T) {
	point := Point{y: 4, x: 9}
	if point == (Point{}) {
		t.Error("Error ", point)
	}

	point = Point{}
	if point != (Point{}) {
		t.Error("Error ", point)
	}
}

//If the type is simple, Golang makes equatable and hashable it
func TestStructEqualForSimpleType(t *testing.T) {
	point := Point{y: 4, x: 9}
	otherPoint := Point{y: 4, x: 9}
	if point != otherPoint {
		t.Error("Error ", point, otherPoint)
	}
}

func TestStructNotEqualForComplexType(t *testing.T) {
	circle := Circle{&Point{y: 4, x: 9}, 3.3}
	otherCircle := Circle{&Point{y: 4, x: 9}, 3.3}

	//Centers have different pointers
	if circle == otherCircle {
		t.Error("Error, ", circle, otherCircle)
	}

	//centers have the same pointers
	otherCircle = Circle{circle.center, 3.3}
	if circle != otherCircle {
		t.Error("Error, ", circle, otherCircle)
	}
}

func TestStructEqualHashKeyMap(t *testing.T) {
	myMap := map[Point]bool{
		Point{y: 4, x: 9}: true,
		Point{y: 4, x: 9}: true,
		Point{y: 4, x: 9}: true,
	}
	if len(myMap) != 1 {
		t.Error("Error ", myMap)
	}
}

type dynamicPoint struct {
	x int
	y int
	z []int
}

func (d dynamicPoint) equal(other dynamicPoint) bool {
	return d.x == other.x && d.y == other.y && len(d.z) == len(other.z)
}

func TestStructEqualForDynamicFields(t *testing.T) {
	point := dynamicPoint{y: 4, x: 9, z: nil}
	otherPoint := dynamicPoint{y: 4, x: 9, z: nil}

	// does not compile
	//if point != otherPoint {}

	if len(point.z) != len(otherPoint.z) {
		t.Error("Error ", point, otherPoint)
	}

	// I need a custom equal method
	if !point.equal(otherPoint) {
		t.Error("Error ", point, otherPoint)
	}
}

type showable interface {
	show() string
}

type door string

func (d door) show() string {
	return "door-name"
}

type window string

func (d window) show() string {
	return "window-name"
}

func TestInterfaceEqual(t *testing.T) {
	var door showable = door("thing")
	var window1 showable = window("thing")

	// they are the same interface but different struct type
	if door == window1 {
		t.Error("Error ", door, window1)
	}

	var window2 showable = window("thing")
	// they are the same interface and struct type
	if window1 != window2 {
		t.Error("Error ", window1, window2)
	}
}

type things struct {
	s []string
}

func (d things) show() string {
	return "things-name"
}

func TestInterfaceWhenHasDynamicFieldShouldBePanic(t *testing.T) {
	t.SkipNow()
	var things1 showable = things{}
	var things2 showable = things{}

	// it compile ok, but it will panic
	if things1 != things2 {
		t.Error("Error ", things1, things2)
	}
}
