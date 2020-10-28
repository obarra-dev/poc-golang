package main

import (
	"testing"
)

func TestNew(t *testing.T) {
	point := new(Point)
	point.x = 10
	point.y = 11
	if point.x == 10 && point.y == 11 {
		t.Log("OK")
	} else {
		t.Error("Error")
	}
}
