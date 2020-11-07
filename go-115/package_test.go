package main_test

import (
	"poc-golang/go-115/mypackage"
	"testing"
)

func TestPackage(t *testing.T) {
	result := mypackage.Language
	if result != "Golang 1.15" {
		t.Error("test Error", result)
	}
}
