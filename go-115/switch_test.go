package main_test

import (
	"testing"
)

func TestSwitch(t *testing.T) {
	environment := "test"
	result := ""
	switch environment {
	case "test", "dev":
		result = "no prod"
	case "prod":
		result = "prod"
	default:
		result = "no case"
	}

	if result == "no prod" {
		t.Log("OK")
	} else {
		t.Error("Error")
	}
}

func TestSwitchWithShortStatement(t *testing.T) {
	result := ""
	switch environment := "test"; environment {
	case "test", "dev":
		result = "no prod"
	case "prod":
		result = "prod"
	default:
		result = "no case"
	}

	if result == "no prod" {
		t.Log("OK")
	} else {
		t.Error("Error")
	}
}

func TestSwitchWithNoCondition(t *testing.T) {
	//This construct can be a clean way to write long if-then-else chains
	environment := "test"
	result := ""
	switch {
	case environment == "test", environment == "dev":
		result = "no prod"
	case environment == "prod":
		result = "prod"
	default:
		result = "no case"
	}

	if result == "no prod" {
		t.Log("OK")
	} else {
		t.Error("Error")
	}
}
