package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

//Types String
// Raw string literals
// Interpreted string literals
func TestStringSplit(t *testing.T) {
	myString := "/api/products/4/cbus/products/other"

	segments := strings.Split(myString, "/products")

	if len(segments) != 3 {
		t.Error("Error", segments, len(segments))
	}
}

func TestStringStrconv(t *testing.T) {
	number, _ := strconv.Atoi("998")

	if number == 998 {
		t.Log("OK")
	} else {
		t.Error("Error", number)
	}
}
func TestIntToString(t *testing.T) {
	result := strconv.Itoa(404)
	if result != "404" {
		t.Error("Error", result)
	}
}

func TestStringsJoin(t *testing.T) {
	array := [3]string{"a", "b", "c"}
	result := strings.Join(array[:], ",")

	if result != "a,b,c" {
		t.Error("Error", result)
	}
}

func TestStringBuilder(t *testing.T) {
	var builder strings.Builder
	builder.WriteString("First")
	builder.WriteString("Second")
	result := builder.String()

	if result != "FirstSecond" {
		t.Error("Error", result)
	}
}

func TestRuneJoinWithStringsBuilder(t *testing.T) {
	array := [3]rune{'a', 'b', 'c'}

	var sb strings.Builder

	for _, char := range array {
		sb.WriteRune(char)
	}

	result := sb.String()

	if result != "abc" {
		t.Error("Error", result)
	}
}

func TestStringIterationWithRange(t *testing.T) {
	namestring := "omar barra"

	var sb strings.Builder
	for _, v := range namestring {

		sb.WriteRune(v)
	}

	result := sb.String()
	if result != "omar barra" {
		t.Error("Error", result)
	}
}

func TestStringIteration(t *testing.T) {
	namestring := "omar barra"

	var sb strings.Builder
	for i := 0; i < len(namestring); i++ {
		sb.WriteByte(namestring[i])
	}

	result := sb.String()
	if result != "omar barra" {
		t.Error("Error", result)
	}
}

//string.Builder also implements the io.Writer interface
func TestStringBuilderFormat(t *testing.T) {
	var sb strings.Builder
	for _, v := range [3]int{1, 2, 3} {
		fmt.Fprintf(&sb, "%d...", v)
	}

	result := sb.String()

	if result != "1...2...3..." {
		t.Error("Error", result)
	}
}

//To write multiline string in GO you can use a raw string literal, where the string is delimited by back quotes
// Raw string literals, delimited by backticks (back quotes), are interpreted literally. They can contain line breaks, and backslashes have no special meaning.
func TestStringMultiLine(t *testing.T) {
	result :=
		`First line
		Second line`

	expected :=
		"First line\n" +
			"Second line"

	if result != expected {
		t.Error("Error", result, expected)
	}
}

func TestAscii(t *testing.T) {
	result := string('o')
	//results := string(11)

	if result != "os" {
		t.Error("Error", result)
	}
}

//both byte and rune are essentially integers
//byte is aliases for unit8
//rune is aliases for int32
func TestBtyeVsRune(t *testing.T) {
	letterO := 'o'
	var letterM byte = 'm'
	var letterA rune = 'a'
	var letterR byte = 114
	letterB := byte('B')

	resultO := string(letterO)
	resultM := string(letterM)
	resultA := string(letterA)
	resultR := string(letterR)
	resultB := string(letterB)

	if resultO != "o" || resultM != "m" ||
		resultA != "a" || resultR != "r" || resultB != "B" {
		t.Error("Error", resultO, resultM, resultA, resultR, resultB)
	}
}

func TestStringToByteSliceAndViceVersa(t *testing.T) {
	nameString := "omar barra"

	nameSlice := []byte(nameString)

	if len(nameSlice) != 10 {
		t.Error("Error", nameSlice)
	}

	other := string(nameSlice)

	if other != nameString {
		t.Error("Error", other, nameString)
	}
}
