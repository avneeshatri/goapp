package main

import (
	"fmt"
	"time"
)

// go run -race to check for race condition.
func main() {
	msg := "Hello"

	go func(msg string) {
		fmt.Println(msg)
	}(msg)
	go func() {
		fmt.Println(msg)
	}()
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}
