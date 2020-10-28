package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	err := fmt.Errorf("Name is invalid.")
	if err != nil {
		t.Log("My error: ", err)
	} else {
		t.Fail()
	}
}

func TestErrorNew(t *testing.T) {
	var ErrorInvalidName = errors.New("Name is invalid.")
	err := ErrorInvalidName
	if err != nil {
		t.Log("My error: ", err)
	} else {
		t.Fail()
	}
}

func TestErrorCustomError(t *testing.T) {
	_, err := executeCustomError()

	if err != nil {
		t.Log("My error: ", err)
	} else {
		t.Fail()
	}
}
