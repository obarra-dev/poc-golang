package main_test

import (
	"fmt"
	"testing"
)

//It is as a WHILE
func TestForInitAndPostStatementsOptional(t *testing.T) {
	var counter int
	for counter <= 10 {
		fmt.Println(counter)
		counter += 1
	}
	if counter == 11 {
		t.Log("OK")
	} else {
		t.Error("Error", counter)
	}
}

//Is as While
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
	}
}

func TestForWithBreakAndWithOutLebel(t *testing.T) {
	var firstCounter int
	var secondCounter int
	for z := 0; z <= 10; z++ {
		for x := 0; x <= 10; x++ {
			if x%2 == 0 {
				break
			}

			firstCounter += x
		}

		secondCounter += z
	}

	fmt.Print(secondCounter)
	if firstCounter == 0 && secondCounter == 55 {
		t.Log("OK")
	} else {
		t.Error("Error", secondCounter)
	}
}

//break with label  terminates other outer (enclosing) loops
func TestForWithBreakAndWithLebel(t *testing.T) {
	var firstCounter int
	var secondCounter int
OuterLoop:
	for z := 0; z <= 10; z++ {
		for x := 0; x <= 10; x++ {
			if x == 3 {
				break OuterLoop
			}

			firstCounter += x
		}

		secondCounter += z
	}

	if firstCounter != 2 && secondCounter != 0 {
		t.Error("Error", firstCounter, secondCounter)
	}
}
