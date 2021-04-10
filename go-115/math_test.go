package main

import (
	"math"
	"testing"
)

func TestNaN(t *testing.T) {
	result := math.NaN()
	if !math.IsNaN(result) {
		t.Error("Error", result)
	}

	result = 0.0
	if math.IsNaN(result) {
		t.Error("Error", result)
	}

	result = 0.1
	if math.IsNaN(result) {
		t.Error("Error", result)
	}
}
