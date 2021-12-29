package main

import (
	"fmt"
	"strconv"
)

const (
	_ = iota + 5
	h = iota + 2
	j = iota
	k
)

func main() {
	var i bool

	fmt.Printf("%v %T\n", i, i)

	n := 1 == 1
	m := 1 == 2

	fmt.Printf(strconv.FormatBool(n) + "\n")
	fmt.Printf(strconv.FormatBool(m) + "\n")

	var o complex128 = 1 + 2i
	fmt.Printf("%v ,%T\n", o, o)
	fmt.Printf("%v ,%T\n", real(o), real(o))
	fmt.Printf("%v ,%T\n", imag(o), imag(o))

	const a = 45
	var b int16 = 34
	fmt.Printf("%v ,%T\n", a+b, a+b)

	fmt.Println(h)
	fmt.Println(j)
	fmt.Println(k)
}
