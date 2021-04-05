package main

import (
	"flag"
	"fmt"
	"testing"
)

// run flag.Parace() is mandatory
func TestFlag(t *testing.T) {
	var consul_URL string
	var defaultVal = "http://consul.local:8500"
	flag.StringVar(&consul_URL, "consul-url", defaultVal, "service discovery url")
	flag.Parse()
	result := getConfig(flag.CommandLine)

	if len(result) != 24 && result[0] != defaultVal {
		t.Error("Error", result)
	}
}

func getConfig(fs *flag.FlagSet) []string {
	cfg := make([]string, 0, 10)
	fs.VisitAll(func(f *flag.Flag) {
		cfg = append(cfg, fmt.Sprintf("%s:%q", f.Name, f.Value.String()))
	})

	return cfg
}
