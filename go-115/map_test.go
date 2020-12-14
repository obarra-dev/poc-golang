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
func TestMapGet(t *testing.T) {
	if myMap["java"] == 1 {
		t.Log("OK")
	} else {
		t.Error("Error")
	}
}

func TestMapGetWithExistingFlag(t *testing.T) {
	if value, ok := myMap["java"]; ok {
		t.Log(value)
	} else {
		t.Error("Error", ok)
	}
}

func TestMapWhenKeyDoesNotExist(t *testing.T) {
	if myMap["kotlin"] == 0 {
		t.Log("OK")
	} else {
		t.Error("Error")
	}
}

func TestMapPut(t *testing.T) {
	myMap["kotlin"] = 10
	if myMap["kotlin"] == 10 {
		t.Log("OK")
	} else {
		t.Error("Error")
	}
}

func TestMapUpdate(t *testing.T) {
	oldValue := myMap["java"]
	myMap["java"] = 10
	if oldValue != 1 || myMap["java"] != 10 {
		t.Error("Error", oldValue, myMap["java"])
	}
}

//if key does not exist, it does return nil??
//If key is not in the map, then elemt is the zero value for the map's element type.
func TestMapDelete(t *testing.T) {
	delete(myMap, "js")
	val, ok := myMap["js"]
	fmt.Println(val)

	if ok == false && val == 0 {
		t.Log("OK")
	} else {
		t.Error("Error")
	}
}

func TestMapIteration(t *testing.T) {
	for key, value := range myMap {
		if value == 0 {
			t.Error("Error: ", key, value)
		}
	}
}
