package main

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

//Types String
// Raw string literals
// Interpreted string literals

func TestStringAreASliceOfBytes(t *testing.T) {
	myStringHex := "\x47\x6f\x20\x69\x73\x20\x41\x77\x65\x73\x6f\x6d\x65\x21"
	if myStringHex != "Go is Awesome!" {
		t.Error("Error ", myStringHex)
	}
	myString := "Go is Awesome!"
	if myString != "Go is Awesome!" {
		t.Error("Error ", myString)
	}

	byte := myStringHex[0]
	if byte != 71 {
		t.Error("Error ", byte)
	}
	byte = myString[0]
	if byte != 71 {
		t.Error("Error ", byte)
	}

	hex := fmt.Sprintf("%x", myStringHex[0])
	if hex != "47" {
		t.Error("Error ", hex)
	}
	hex = fmt.Sprintf("%x", myString[0])
	if hex != "47" {
		t.Error("Error ", hex)
	}

	char := fmt.Sprintf("%q", myStringHex[0])
	if char != "'G'" {
		t.Error("Error ", char)
	}
	char = fmt.Sprintf("%q", myString[0])
	if char != "'G'" {
		t.Error("Error ", char)
	}
}

func TestStringTrim(t *testing.T) {
	myString := "    never Trust a Programmer    "
	r := strings.Trim(myString, " ")
	if r != "never Trust a Programmer" {
		t.Error("Error", r, len(r))
	}
	r = strings.TrimSpace(myString)
	if r != "never Trust a Programmer" {
		t.Error("Error", r, len(r))
	}

	myString = "    77never Trust a Programmer!!!    "
	r = strings.TrimFunc(myString, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	if r != "never Trust a Programmer" {
		t.Error("Error", r, len(r))
	}

	myString = "    never Trust a Programmer    "
	if strings.TrimLeft(myString, " ") != "never Trust a Programmer    " ||
		strings.TrimRight(myString, " ") != "    never Trust a Programmer" {
		t.Error("Error")
	}
	if strings.TrimPrefix(myString, "    ") != "never Trust a Programmer    " ||
		strings.TrimSuffix(myString, "    ") != "    never Trust a Programmer" {
		t.Error("Error")
	}
}

func TestStringContains(t *testing.T) {
	myString := "never Trust a Programmer"
	r := strings.Contains(myString, "Programmer")
	if !r {
		t.Error("Error", r)
	}

	if !strings.HasPrefix(myString, "never") || !strings.HasSuffix(myString, "Programmer") {
		t.Error("Error")
	}
}

func TestStringRemplace(t *testing.T) {
	myString := "I am Programmer but never Trust a Programmer"
	r := strings.Replace(myString, "Programmer", "Lawer", -1)
	if r != "I am Lawer but never Trust a Lawer" {
		t.Error("Error", r)
	}

	r = strings.Replace(myString, "Programmer", "Lawer", 1)
	if r != "I am Lawer but never Trust a Programmer" {
		t.Error("Error", r)
	}

	r = strings.ReplaceAll(myString, "Programmer", "Lawer")
	if r != "I am Lawer but never Trust a Lawer" {
		t.Error("Error", r)
	}
}

// compare is faster than equal operator, but for Basic comparison equal are somtimes faster
func TestStringCompare(t *testing.T) {
	r := strings.Compare("Omar", "omar")
	if r != -1 {
		t.Error("Error ", r)
	}

	r = strings.Compare("omar", "Omar")
	if r != 1 {
		t.Error("Error ", r)
	}

	//it is case sensitive
	r = strings.Compare("omar", "omar")
	if r != 0 {
		t.Error("Error ", r)
	}
}

func TestStringSubstring(t *testing.T) {
	myString := "/api/products/4/cbus/products/other"
	segment := myString[3:6]
	if segment != "i/p" {
		t.Error("Error", segment)
	}
}

func TestStringSplit(t *testing.T) {
	myString := "/api/products/4/cbus/products/other"

	segments := strings.Split(myString, "/")
	if len(segments) != 7 || segments[0] != "" || segments[1] != "api" {
		t.Error("Error", segments, len(segments))
	}

	segments = strings.SplitAfter(myString, "/")
	if len(segments) != 7 || segments[0] != "/" || segments[1] != "api/" {
		t.Error("Error", segments, len(segments))
	}

	segments = strings.SplitAfterN(myString, "/", 2)
	if len(segments) != 2 || segments[0] != "/" || segments[1] != "api/products/4/cbus/products/other" {
		t.Error("Error", segments, len(segments))
	}
}

func TestStringFields(t *testing.T) {
	myString := "never Trust a Programmer"
	r := strings.Fields(myString)
	if len(r) != 4 || r[0] != "never" || r[1] != "Trust" {
		t.Error("Error", r, len(r))
	}
}

func TestStringFieldsFunc(t *testing.T) {
	f := func(c rune) bool {
		return unicode.IsSpace(c) || c == '.'
	}
	myString := "never Trust a Programmer. Never."
	r := strings.FieldsFunc(myString, f)
	if len(r) != 5 || r[0] != "never" || r[1] != "Trust" {
		t.Error("Error", r, len(r))
	}
}

func TestStringCasing(t *testing.T) {
	myString := "never Trust a Programmer"
	r := strings.ToLower(myString)
	if r != "never trust a programmer" {
		t.Error("Error", r, myString)
	}

	r = strings.ToUpper(myString)
	if r != "NEVER TRUST A PROGRAMMER" {
		t.Error("Error", r, myString)
	}
	r = strings.Title(myString)
	if r != "Never Trust A Programmer" {
		t.Error("Error", r, myString)
	}
}

func TestStringRepeat(t *testing.T) {
	myString := "Programmer"
	r := strings.Repeat(myString, 5)
	if r != "ProgrammerProgrammerProgrammerProgrammerProgrammer" {
		t.Error("Error", r)
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
	if result != "o" {
		t.Error("Error", result)
	}

	var letterR byte = 114
	result = string(letterR)
	if result != "r" {
		t.Error("Error", result)
	}
}

//both byte and rune are essentially integers
//byte is aliases for unit8
//rune is aliases for int32
func TestByteVsRune(t *testing.T) {
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
