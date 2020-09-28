package main

import (
	"testing"
)

func TestPolymorphism(t *testing.T) {
	circle := Circle{&Point{y: 4, x: 9}, 3.3}
	circleFlatten := CircleFlatten{&Point{6, 7}, 5.5}

	shapes := []shape{&circle, &circleFlatten}
	for _, shape := range shapes {
		shape.setDoubleRadius(2.0)
		result := shape.getDoubleRadius()
		if result == 4.0 {
			t.Log("test ok")
		} else {
			t.Error("test Error")
			t.Fail()
		}
	}
}
