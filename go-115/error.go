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

type omarError struct {
	code    string
	message string
}

func (m *omarError) Error() string {
	return fmt.Sprintf("Omar Error: %s - %s", m.code, m.message)
}

func NewOmarError(code string, message string) error {
	return &omarError{code, message}
}

type barraError struct {
	query string
	err   error
}

func (e *barraError) Error() string {
	return fmt.Sprintf("Barra Error: %s - %s", e.query, e.err.Error())
}

func (e *barraError) Unwrap() error {
	return e.err
}

func NewBarraError(query string, err error) error {
	return &barraError{query, err}
}
