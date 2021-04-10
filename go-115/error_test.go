package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrorExamingError(t *testing.T) {
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

// adds information but discards everything from the original error except the text
func TestErrorAddingInformation(t *testing.T) {
	err := fmt.Errorf("reading %v: %v", "File", errors.New("name is invalid"))
	if err.Error() != "reading File: name is invalid" {
		t.Error("Error", err)
	}
}

//Wrap error
func TestErroUnwraprWrapOtherError(t *testing.T) {
	err := fmt.Errorf("reading %v: %w", "File", errors.New("name is invalid"))
	if err.Error() != "reading File: name is invalid" {
		t.Error("Error", err)
	}

	errWrapped := errors.Unwrap(err)
	if errWrapped.Error() != "name is invalid" {
		t.Error("Error", errWrapped)
	}
}

// sentinel value
var ErrorInvalidName = errors.New("name is invalid")

func TestErrorCompareSentinelValueBefore113(t *testing.T) {
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

//The errors.Is function compares an error to a value.
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

func TestErrorIsSentinelValueAfter113WithWrapper(t *testing.T) {
	omarErrSentinel := NewOmarError("04", "error validating")

	barraErr := NewBarraError("Todo Mal", omarErrSentinel)
	//the omarErrSentinel instance is in barraErr
	if !errors.Is(barraErr, omarErrSentinel) {
		t.Error("Error", barraErr)
	}

	barraErr = NewBarraError("Todo Mal", NewOmarError("04", "error validating"))
	//the omarErrSentinel instance is not in barraErr
	if errors.Is(barraErr, omarErrSentinel) {
		t.Error("Error", barraErr)
	}

	barraErr = NewBarraError("Todo Mal", omarErrSentinel)
	otherErr := validateName("omar")
	//the  some other type instance is no in barraErr
	if errors.Is(barraErr, otherErr) {
		t.Error("Error", otherErr)
	}

	barraErr = NewBarraError("Todo Mal", otherErr)
	//the some other type sentinel err instance is in barraErr, but it is no a omarErrSentinel
	if errors.Is(barraErr, omarErrSentinel) {
		t.Error("Error", barraErr)
	}
}

func TestErrorCustomError(t *testing.T) {
	err := NewCustomError("04", "error validating")
	if err == nil || err.Error() != "Custom Error: 04 - error validating" {
		t.Error("Error generating custom error", err)
	}
}

func TestErrorAssertinTypeBefore113(t *testing.T) {
	err := NewOmarError("04", "error validating")
	if e, ok := err.(*omarError); !ok {
		t.Error("Error ", e, ok)
	}

	err = NewCustomError("04", "error validating")
	if _, ok := err.(*omarError); ok {
		t.Error("Error ", ok)
	}
}

//The As function tests whether an error is a specific type.
func TestErrorAssertinTypeAfter113(t *testing.T) {
	err := NewOmarError("04", "error validating")
	var omarError *omarError
	if !errors.As(err, &omarError) {
		t.Error("Error ", omarError)
	}

	err = NewCustomError("04", "error validating")
	if errors.As(err, &omarError) {
		t.Error("Error ", omarError)
	}
}

func TestErrorAssertinTypeAfter113Wrapper(t *testing.T) {
	var omarError *omarError
	omarErrSentinel := NewOmarError("04", "error validating")

	barraErr := NewBarraError("Todo Mal", omarErrSentinel)
	//some instance of omarError is in barraErr
	if !errors.As(barraErr, &omarError) {
		t.Error("Error", barraErr)
	}

	barraErr = NewBarraError("Todo Mal", NewOmarError("04", "error validating"))
	//some instance of omarError is in barraErr
	if !errors.As(barraErr, &omarError) {
		t.Error("Error", barraErr)
	}

	barraErr = NewBarraError("Todo Mal", omarErrSentinel)
	var customError *customError
	if errors.As(barraErr, &customError) {
		t.Error("Error", barraErr)
	}

	barraErr = NewBarraError("Todo Mal", NewCustomError("04", "error validating"))
	if errors.As(barraErr, &omarError) {
		t.Error("Error", barraErr)
	}

}

func validateName(name string) error {
	return ErrorInvalidName
}
