package main_test

import (
	"errors"
	"io"
	"testing"
	"time"
)

func TestConstDouble(t *testing.T) {
	const PI = 3.14
	//can not assign
	//PI = 3.5
	if PI != 3.14 {
		t.Error("test Error")
	}
}

func TestConstString(t *testing.T) {
	const PI = "3.14"
	if PI != "3.14" {
		t.Error("test Error")
	}
}

//Go' constant are powerful. If you only think of them as immutable numbers, you' re missing out

//constant whose value is the number of bits in machine word
//constant expression happen in compile time
func TestConstExpression(t *testing.T) {
	const intSize = 32 << (^int(0) >> 32 & 1)
	if intSize != 64 {
		t.Error("Error ", intSize)
	}
	const one = 1e3 - 99.0*10.0 - 9
	if one != 1 {
		t.Error("Error ", one)
	}

}

//Constant values can implement interfaces
func TestConstImplementingInterface(t *testing.T) {
	const timeout = 500 * time.Millisecond
	var expected time.Duration = 500000000

	if timeout.String() != "500ms" || timeout != expected {
		t.Error("Error", timeout, expected)
	}
}

//Sentinel errors are exported public variables
func TestSentinalValueAreNotConstant(t *testing.T) {
	if io.EOF.Error() != "EOF" {
		t.Error("Error", io.EOF)
	}

	//sentinel error are not inmutable so it can change
	io.EOF = errors.New("new error")
	if io.EOF.Error() != "new error" {
		t.Error("Error", io.EOF)
	}
}

type MyError string

func (e MyError) Error() string {
	return string(e)
}

func TestConstCreateInmutableError(t *testing.T) {
	const myError = MyError("CustomError")
	const myError2 = MyError("CustomError")

	if myError != myError2 {
		t.Error("Error ", myError2, myError)
	}
}

const (
	myConst1 = iota / 0.5
	myConst2
	myConst3
	myConst4
	myConst5
)

func TestConstIota(t *testing.T) {
	if myConst1 != 0 || myConst2 != 2 || myConst3 != 4 || myConst4 != 6 || myConst5 != 8 {
		t.Error("Error ", myConst1, myConst2, myConst3, myConst4, myConst5)
	}
}
