package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func executeWritersAndReadersRWMutex() {
	myMap := map[int]int{}

	mux := &sync.RWMutex{}

	go writeLoop(myMap, mux)
	go readLoop(myMap, mux)
	go readLoop(myMap, mux)
	go readLoop(myMap, mux)
	go readLoop(myMap, mux)

	mytimeOutChan := time.After(2 * time.Second)

	for {
		select {
		case <-mytimeOutChan:
			fmt.Println("TIME OUT ")
			time.Sleep(1 * time.Second)
			os.Exit(0)
		}

	}
}

func writeLoop(myMap map[int]int, rwmux *sync.RWMutex) {
	for {
		for i := 0; i < 100; i++ {
			rwmux.Lock()
			myMap[i] = i
			rwmux.Unlock()
		}
	}
}

func readLoop(myMap map[int]int, rwmux *sync.RWMutex) {
	for {
		rwmux.RLock()
		for k, v := range myMap {
			fmt.Println(k, "-", v)
		}
		rwmux.RUnlock()
	}
}
