package main

import (
	"encoding/json"
	"testing"
)

//The DeeplyEqual() method is defined under “reflect” package.

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

// Embedding type See embedding_test.go
func TestStructEmbedded(t *testing.T) {
	circle := CircleFlatten{&Point{4, 9}, 3.3}
	circle.x = 10
	if circle.x != 10 || circle.y != 9 || circle.redius != 3.3 {
		t.Error("Error", circle.x, circle.y)
	}
}

func TestAnonymousStructs(t *testing.T) {
	var config struct {
		APIKey string
		ID     int
	}
	config.APIKey = "BADC0C0A"
	config.ID = 10

	c := struct {
		APIKey string
		ID     int
	}{
		APIKey: "BADC0C0A",
		ID:     10,
	}

	if config != c {
		t.Error("Error ", c, config)
	}
}

//Cheaper and safer than using map[string]interface{}
func TestAnonymousStructs2(t *testing.T) {
	//OK case
	b := []byte(`{"APIKey":"Maru","ID":123}`)

	c := struct {
		APIKey string
		ID     int
	}{}

	err := json.Unmarshal(b, &c)

	if err != nil {
		t.Error("Error ", err)
	}

	if c.APIKey != "Maru" {
		t.Error("Error", c)
	}

	//ERROR case
	b = []byte(`{"APIKey":123,"ID":123}`)
	err = json.Unmarshal(b, &c)
	// with anonymous structs you can check type
	if err == nil {
		t.Error("Error ", err)
	}
}

func TestAnonymousStructs2UsingMapInterface(t *testing.T) {
	//OK case
	b := []byte(`{"APIKey":"Maru","ID":123}`)

	c := map[string]interface{}{}

	err := json.Unmarshal(b, &c)
	if err != nil {
		t.Error("Error unmarshal ", err)
	}

	key, ok := c["APIKey"]
	if !ok {
		t.Error("Error ", c)
	}

	keyString, ok := key.(string)
	if !ok {
		t.Error("Error ", c)
	}

	if keyString != "Maru" {
		t.Error("Error", c)
	}

	//ERROR case
	c = map[string]interface{}{}
	b = []byte(`{"APIKey":123,"ID":123}`)
	err = json.Unmarshal(b, &c)
	// with map interface  you can not check type
	if err != nil {
		t.Error("Error ", err)
	}

	key, ok = c["APIKey"]
	if !ok {
		t.Error("Error ", c)
	}

	//you need extra work to check type
	keyString, ok = key.(string)
	if ok {
		t.Error("Error ", c, keyString)
	}

}
