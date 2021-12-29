package main

import (
	"fmt"
)

type Car struct {
	model  string
	number int
}

func main() {
	var a map[string]int = map[string]int{"california": 23, "texas": 45}
	fmt.Println(a)
	b := map[string]int{"x": 2, "y": 1}
	fmt.Println(b)
	c := b
	c["x"] = 4
	fmt.Println(b)

	m := make(map[string]int, 10)
	m = map[string]int{"x": 12}

	fmt.Println(m["x"])
	fmt.Println(len(m))
	delete(a, "texas")
	fmt.Println(a)

	val, ok := a["dummy"]
	fmt.Println(val, ok)

	toyota := Car{number: 1234, model: "Toyota"}
	fmt.Println(toyota)

	doctor := struct{ name string }{name: "Sharma"}
	fmt.Println(doctor)

	doctor2 := &doctor
	doctor2.name = "Mehta"

	fmt.Println(doctor)
	fmt.Println(doctor2)
}
