package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(1)
	go sayHello()

	msg := "Hi there"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "Hey"
	wg.Wait()
}

func sayHello() {
	fmt.Println("Hello")
	wg.Done()
}
