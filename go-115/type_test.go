package main_test

import (
	"strconv"
	"testing"
)

//Type aliases
type day = int

//it does not compile, becaouse day can not be extended
//func (d day) doSomethingDay() {}

//day and int are the same
func TestTypeAlias(t *testing.T) {
	var n int = 10
	var a day = n
	if a != 10 {
		t.Error("test Error", a)
	}
}

//Type definition
type month int

//it compile, becaouse day can be extended
func (d month) doSomethingMonth() month { return d }

//The new type is called a defined type. It is different from any other type, including the type it is created from.
func TestTypeDefinition(t *testing.T) {
	var n int = 10
	//var a month = n  //compile error

	var m month = month(n)
	if m != 10 {
		t.Error("test Error", m)
	}

	// can not assign other type but you can assign the orginal type
	m = 15
	if m != 15 {
		t.Error("test Error", m)
	}

	m = m.doSomethingMonth()
	if m != 15 {
		t.Error("test Error", m)
	}
}

// it is possible because,  it will have the same memory layout and struct,
func TestTypeConversion(t *testing.T) {
	// var m month = 10 // it compile!!, it is ok
	var m month = month(10)
	if m != 10 {
		t.Error("test Error", m)
	}

	// var b int = int(m) // it no compile!!
	var b int = int(m)
	if b != 10 {
		t.Error("test Error", b)
	}
}

type LeaveType string

const (
	AnnualLeave LeaveType = "AnnualLeave"
	Sick                  = "Sick"
	BankHoliday           = "BankHoliday"
	Other                 = "Other"
)

func TestTypeDefinitionConversionWithInvalidValue(t *testing.T) {
	// it compiles, so TypeAlias allows any values
	var lt LeaveType = LeaveType("Hello")
	if lt != "Hello" {
		t.Error("test Error", lt)
	}
}

//Type definition
type app []string

func TestTypeDefinition2(t *testing.T) {
	apps := app{"Facebook", "Instagram", "WhatsApp"}
	if len(apps) != 3 {
		t.Error("test Error", apps)
	}
}

type Vehicle int

const (
	Bike Vehicle = iota
	Bus
	Train
)

func TestTypeDefinitionConstantsUsingIota(t *testing.T) {
	if Bike != 0 || Bus != 1 || Train != 2 {
		t.Error("test Error", Bike, Bus, Train)
	}
}

func TestTypeDefinitionForMap(t *testing.T) {
	integers := []int{1, 2, 3}
	strings := Map(func(x int) string { return strconv.Itoa(x) }, integers)
	if len(strings) != 3 || strings[0] != "1" {
		t.Error("test Error", strings)
	}
}

type operatorIntString func(x int) string

func Map(op operatorIntString, a []int) []string {
	res := make([]string, len(a))
	for i, x := range a {
		res[i] = op(x)
	}
	return res
}
