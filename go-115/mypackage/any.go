package mypackage

import (
	"fmt"
)

//Package initialization is done only once even if package is imported many times.
func init() {
	fmt.Println("Init function 1")
}

//Many init functions can be defined in the same package or file
func init() {
	fmt.Println("Init function 2")
}

const (
	Language = "Golang 1.15"
)
