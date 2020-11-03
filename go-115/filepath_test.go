package main_test

import (
	"path/filepath"
	"testing"
)

func TestFilePath(t *testing.T) {
	path := filepath.Join("dirOne", "dirTwo")

	if path == "dirOne\\dirTwo" {
		t.Log("OK")
	} else {
		t.Error("Error", path)
	}
}
