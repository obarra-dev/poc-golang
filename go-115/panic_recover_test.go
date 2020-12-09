package main

import (
	"fmt"
	"testing"
)

//is a built-in function that stops the ordinary flow of control and begins panicking.
//The process continues up the stack until all functions in the current goroutine have returned
//So you can see the stacktrace
func TestPanic(t *testing.T) {
	t.SkipNow()
	doSomethingWithOutRecover(1)
}

func TestRecover(t *testing.T) {
	result := doSomethingWithRecover()
	fmt.Println("Recovered")
	if result == "Returned normally from doSomethingDeferRecover." {
		t.Log("OK")
	} else {
		t.Error("Error", result)
	}
}

func doSomethingWithOutRecover(counter int) {
	counterAndDoPanic(counter)
}
func counterAndDoPanic(counter int) {
	if counter > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", counter))
	}
	counterAndDoPanic(counter + 1)
}

func doSomethingWithRecover() string {
	doSomethingDeferRecover()
	return "Returned normally from doSomethingDeferRecover."
}

func doSomethingDeferRecover() {
	fmt.Printf("Doing something...")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in doSomethingDeferRecover:", r)
		} else {
			fmt.Println("Error recovering!!")
		}

	}()
	counterAndDoPanic(1)
	fmt.Println("Returned normally from counterWithPanic. Never happen")
}
