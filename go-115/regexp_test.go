package main

import (
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	myString := "This is a string. There are many strings lit it, but his is mine"
	regex, _ := regexp.Compile(`s(\w[a-z]+)g\b`)

	if !regex.MatchString(myString) {
		t.Error("Error ", regex.MatchString(myString))
	}

	matches := regex.FindAllString(myString, -1)
	if len(matches) != 1 || matches[0] != "string" {
		t.Error("Error ", len(matches))
	}

	loc := regex.FindStringIndex(myString)
	if len(loc) != 2 || loc[0] != 10 || loc[1] != 16 {
		t.Error("Error ", len(loc), loc[0])
	}

	replaced := regex.ReplaceAllString(myString, "Omar")
	if replaced != "This is a Omar. There are many strings lit it, but his is mine" {
		t.Error("Error ", replaced)
	}

}
