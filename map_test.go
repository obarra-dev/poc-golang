package main_test

import (
	"fmt"
	"testing"
)

var myMap = map[string]int{
	"golang": 5,
	"js":     3,
	"java":   1,
}

//arrays cannot be resized.
func TestMapGetSimple(t *testing.T) {
	if myMap["java"] == 1 {
		t.Log("OK")
	} else {
		t.Error("Error")
		t.Fail()
	}
}

func TestMapPut(t *testing.T) {
	myMap["kotlin"] = 10
	if myMap["kotlin"] == 10 {
		t.Log("OK")
	} else {
		t.Error("Error")
		t.Fail()
	}
}

func TestMapUpdate(t *testing.T) {
	myMap["java"] = 10
	if myMap["java"] == 10 {
		t.Log("OK")
	} else {
		t.Error("Error")
		t.Fail()
	}
}

//if key does not exist, it does return nil
//If key is not in the map, then elem is the zero value for the map's element type.
func TestMapDelete(t *testing.T) {
	delete(myMap, "js")
	val, ok := myMap["js"]
	fmt.Println(val)

	if ok == false && val == 0 {
		t.Log("OK")
	} else {
		t.Error("Error")
		t.Fail()
	}
}
