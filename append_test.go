package main_test

import (
	"testing"
)

func TestAppendSimple(t *testing.T) {
	var arraWithOutLimit []int
	arraWithOutLimit = append(arraWithOutLimit, 10, 12, 13)
	if len(arraWithOutLimit) == 3 && cap(arraWithOutLimit) == 4 {
		t.Log("OK")
	} else {
		t.Error("Error", len(arraWithOutLimit), cap(arraWithOutLimit))
		t.Fail()
	}
}

func TestAppend(t *testing.T) {
	arraWithOutLimit := []int{3, 4}
	arraWithOutLimit = append(arraWithOutLimit, 10, 12, 13)
	if len(arraWithOutLimit) == 5 && cap(arraWithOutLimit) == 6 {
		t.Log("OK")
	} else {
		t.Error("Error", len(arraWithOutLimit), cap(arraWithOutLimit))
		t.Fail()
	}
}
