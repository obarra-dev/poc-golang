package main_test

import (
	"testing"
)

func TestMake(t *testing.T) {
	maked := make([]int, 5)
	if len(maked) == 5 && cap(maked) == 5 {
		t.Log("OK")
	} else {
		t.Error("Error", len(maked), cap(maked))
	}
}

func TestMakeCapAndLengDiff(t *testing.T) {
	maked := make([]int, 0, 5)
	if len(maked) == 0 && cap(maked) == 5 {
		t.Log("OK")
	} else {
		t.Error("Error", len(maked), cap(maked))
	}
}

func TestMakeMap(t *testing.T) {
	maked := make(map[string]int)

	if len(maked) == 0 {
		t.Log("OK")
	} else {
		t.Error("Error", len(maked))
	}
}
