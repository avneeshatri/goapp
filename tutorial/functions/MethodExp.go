package main

import (
	"fmt"
)

func main() {

	g := greeter{
		greeting: "Hello",
		name:     "Vishu",
	}
	g.greet()
	fmt.Println(g.name)
}

type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "mohit"
}
