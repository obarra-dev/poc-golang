package main

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	doSomething()
}

func TestRecover(t *testing.T) {
	doSomethingWithRecover()
	fmt.Println("Recovered")
}

func doSomething() {
	panic("Panic Test")
}

func doSomethingWithRecover() {
	fmt.Printf("Doing something...")
	defer func() {
		messageError := recover()
		fmt.Println(messageError)
	}()
	panic("Panic Test with recover")
}
