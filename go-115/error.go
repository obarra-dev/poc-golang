package main

import (
	"fmt"
)

type customError struct {
	code    string
	message string
}

func (m *customError) Error() string {
	return fmt.Sprintf("Custom Error: %s - %s", m.code, m.message)
}

func NewCustomError(code string, message string) error {
	return &customError{code, message}
}

type OmarError struct {
	code    string
	message string
}

func (m *OmarError) Error() string {
	return fmt.Sprintf("Omar Error: %s - %s", m.code, m.message)
}

func NewOmarError(code string, message string) error {
	return &OmarError{code, message}
}
