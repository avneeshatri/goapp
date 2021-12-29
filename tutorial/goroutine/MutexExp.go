package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {

	for i := 0; i < 20; i++ {
		wg.Add(2)
		m.RLock()
		go printCounter()
		m.Lock()
		go incrementCounter()

	}

	wg.Wait()
}

func incrementCounter() {
	counter++
	m.Unlock()
	wg.Done()
}

func printCounter() {
	fmt.Println("Counter", counter)
	m.RUnlock()
	wg.Done()
}
