package main

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2, 3, 4}
	a[0] = 23
	fmt.Printf("Array: %v\n", a)

	var b [3]string
	b[2] = "Lisa"
	fmt.Printf("Array: %v\n", b)
	fmt.Printf("Length of b: %v\n", len(b))

	c := &a
	c[1] = 45
	fmt.Printf("Array: %v\n", a)
	fmt.Println(a)
	fmt.Println(c)

	// slice

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	y := x
	y[2] = 45
	fmt.Println(x)
	fmt.Printf("Length of b: %v\n", len(x))
	fmt.Printf("Capacity of b: %v\n", cap(x))

	z := y[3:6]
	fmt.Println(z)

	var l = make([]int, 3, 5)
	fmt.Println(l)
	fmt.Printf("Length of l: %v\n", len(l))
	fmt.Printf("Capacity of l: %v\n", cap(l))
	l = append(l, 4)
	l = append(l, 5)
	l = append(l, 6, 7, 8, 9, 10, 11)

	l = append(l, []int{12, 13, 14, 15}...)

	fmt.Println(l)
	fmt.Printf("Length of l: %v\n", len(l))
	fmt.Printf("Capacity of l: %v\n", cap(l))
}
