package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func ExecuteGorutines() {
	waitGroup.Add(2)

	go showHolder()

	go showInsurer()

	fmt.Println("Waiting Proccess")
	waitGroup.Wait()
	fmt.Println("End Proccess")
}

func showHolder() {
	defer waitGroup.Done()

	for _, letter := range "IamaHolder" {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Holder: ", letter)
	}
}

func showInsurer() {
	defer waitGroup.Done()

	for _, letter := range "IamaInsurer" {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Insurer: ", letter)
	}
}
