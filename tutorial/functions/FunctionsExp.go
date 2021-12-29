package main

import (
	"fmt"
)

func main() {
	greeting := "hi"
	var name string
	name = "suresh"
	greetings(&greeting, &name)
	fmt.Println(name)

	total := sum(1, 2, 3, 4)
	fmt.Println("sum", *total)

	sumNmdRetVal := sumNamedReturnValue(5, 6, 7, 8)
	fmt.Println(sumNmdRetVal)

	result, err := divide(12, 0)

	if err != nil {
		fmt.Println("Error:", err.Error())
		fmt.Println("Division failed")
	} else {
		fmt.Println(result)
	}
}

func divide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}

	return a / b, nil
}

func sum(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

func sumNamedReturnValue(values ...int) (result int) {
	fmt.Println(values)
	result = 0
	for _, v := range values {
		result += v
	}
	return
}

func greetings(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "mahesh"
	fmt.Println(*name)
}
