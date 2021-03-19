package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrorExamingErrorBefore113(t *testing.T) {
	err := errors.New("name is invalid")
	if err != nil {
		t.Log("Examing error before 113")
	} else {
		t.Error("Error err ", err)
	}
}

func TestErrorWhenGenerateError(t *testing.T) {
	errF := fmt.Errorf("name is invalid")
	if errF == nil {
		t.Error("Error generating error with fmt.ErroF")
	}

	errN := errors.New("name is invalid")
	if errN == nil {
		t.Error("Error generating error with errors.New")
	}
}

func TestErrorMethodErrorShouldBeReturnStringValue(t *testing.T) {
	var err = errors.New("name is invalid")
	if err.Error() != "name is invalid" {
		t.Error("Error", err)
	}
}

func TestErrorAddingInformation(t *testing.T) {
	err := fmt.Errorf("reading %v: %v", "File", errors.New("name is invalid"))
	if err.Error() != "reading File: name is invalid" {
		t.Error("Error", err)
	}
}

// sentinel value
var ErrorInvalidName = errors.New("name is invalid")

func TestErrorCompereSentinelValueBefore113(t *testing.T) {
	err := errors.New("name is invalid")
	//the errors has the same value but are different instance
	if err == ErrorInvalidName {
		t.Error("It is imposible", err)
	}

	err = validateName("omar")
	if err != ErrorInvalidName {
		t.Error("The erros has no the same instances", err)
	}
}

func TestErrorIsSentinelValueAfter113(t *testing.T) {
	err := errors.New("name is invalid")
	//the errors has the same value but are different instance
	if errors.Is(err, ErrorInvalidName) {
		t.Error("It is imposible", err)
	}

	err = validateName("omar")
	if !errors.Is(err, ErrorInvalidName) {
		t.Error("The erros has no the same instances", err)
	}
}

func TestErrorCustomError(t *testing.T) {
	err := NewCustomError("04", "error validating")
	if err == nil || err.Error() != "Custom Error: 04 - error validating" {
		t.Error("Error generating custom error", err)
	}
}

func TestErrorUnwrappingBefore113(t *testing.T) {
	err := NewOmarError("04", "error validating")
	if e, ok := err.(*OmarError); !ok {
		t.Error("Error ", e, ok)
	}

	err = NewCustomError("04", "error validating")
	if _, ok := err.(*OmarError); ok {
		t.Error("Error ", ok)
	}
}

func TestErrorUnwrappingAfter113(t *testing.T) {
	err := NewOmarError("04", "error validating")
	var omarError *OmarError
	if !errors.As(err, &omarError) {
		t.Error("Error ", omarError)
	}

	err = NewCustomError("04", "error validating")
	if errors.As(err, &omarError) {
		t.Error("Error ", omarError)
	}

}

func validateName(name string) error {
	return ErrorInvalidName
}
