package main_test

import (
	"fmt"
	"testing"
)

//It is as a WHILE
func TestForInitAndPostStatementsOptional(t *testing.T) {
	var counter int
	for counter <= 10 {
		counter += counter
	}
	if counter == 9 {
		t.Log("OK")
	} else {
		t.Error("Error")
		t.Fail()
	}
}

func TestForForever(t *testing.T) {
	var counter int
	for {
		counter++
		if counter == 5 {
			break
		}
	}
	if counter == 5 {
		t.Log("OK")
	} else {
		t.Error("Error")
		t.Fail()
	}
}

func TestForWithContinue(t *testing.T) {
	var counter int
	for x := 0; x <= 10; x++ {
		if x%2 == 0 {
			continue
		}

		counter += x
	}

	fmt.Print(counter)
	if counter == 25 {
		t.Log("OK")
	} else {
		t.Error("Error")
		t.Fail()
	}
}

func TestForWithBreak(t *testing.T) {
	var counter int
	for x := 0; x <= 10; x++ {
		if x%2 == 0 {
			break
		}

		counter += x
	}

	fmt.Print(counter)
	if counter == 0 {
		t.Log("OK")
	} else {
		t.Error("Error")
		t.Fail()
	}
}
