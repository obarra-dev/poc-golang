package main_test

import "testing"

func TestVariablesfdfd(t *testing.T) {
	integer := 19
	if integer == 19 {
		t.Log("test ok")
	} else {
		t.Error("test Error")
		t.Fail()
	}
}
