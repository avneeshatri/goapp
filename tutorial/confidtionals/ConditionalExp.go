package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("condition pass")
	}

	m := map[string]int{"delhi": 1, "meerut": 2}

	if prop, ok := m["delhi"]; ok {
		fmt.Println(prop)
		fmt.Println(prop)
	} else {
		fmt.Printf("Prop not found")
	}

	if 4 > 5 && returnTrue() {
		fmt.Println("Short circut test")
	}

	i := 5

	switch i {
	case 1, 5, 10:
		fmt.Println("case 1 5 or 10")
	case 2:
		fmt.Println("case 2")
	default:
		fmt.Println("default")
	}

	switch {
	case 1 < 2:
		fmt.Println("case 1")
		fallthrough
	case 3 > 5:
		fmt.Println("case 2")
	case 2 > 5:
		fmt.Println("case 3")
	default:
		fmt.Println("default")
	}

	var j interface{} = 1
	switch j.(type) {
	case int:
		fmt.Println("case int")
		break
		fmt.Println("after break")
	case float64:
		fmt.Println("case float64")
	case bool:
		fmt.Println("case bool")
	default:
		fmt.Println("default")
	}
}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}
