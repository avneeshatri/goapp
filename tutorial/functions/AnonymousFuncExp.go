package main

import (
	"fmt"
)

func main() {
	f := func() {
		fmt.Println("Hello!")
	}
	f()

	var div func(float32, float32) (float32, error)
	div = func(a, b float32) (float32, error) {
		if b == 0 {
			return 0, fmt.Errorf("Cannot divide by zero")
		}

		return a / b, nil
	}

	result, err := div(13, 3)

	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err.Error())
	}
}
