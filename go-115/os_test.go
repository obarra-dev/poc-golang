package main

import (
	"os"
	"testing"
)

func TestOSLookupEnv(t *testing.T) {
	if val, ok := os.LookupEnv("CONSUL_URL"); ok {
		t.Error("Error", val)
	}
}

func TestOSGetenv(t *testing.T) {
	if val := os.Getenv("CONSUL_URL"); val != "" {
		t.Error("Error", val)
	}
}
