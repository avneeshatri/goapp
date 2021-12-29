package main

import (
	"fmt"
)

func main() {
	var writer Writer = &ConsoleWriter{}
	writer.Write([]byte("Hello there"))

	i := IntCounter(0)

	for j := 0; j < 5; j++ {
		fmt.Println(i.increment())

	}
}

type IntCounter int

type Incrementer interface {
	increment() int
}

func (inc *IntCounter) increment() int {
	*inc++
	return int(*inc)
}

type Writer interface {
	Write(data []byte) (int, error)
}

type ConsoleWriter struct {
}

func (c *ConsoleWriter) Write(data []byte) (int, error) {
	fmt.Println(string(data))
	return len(data), nil
}
