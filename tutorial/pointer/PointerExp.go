package main

import (
	"fmt"
)

func main() {
	var a int = 21
	var b *int = &a
	m := map[string]int{"delhi": 1, "meerut": 2}
	n := &m
	c := &a
	*c = 23
	fmt.Println(&a, *b, *c)
	fmt.Println(m, n)

	l := [3]int{1, 2, 3}
	p := &l[1]
	q := &l[2]

	*q = 4

	fmt.Printf("%v %p %p\n", l, p, q)

	var toyota *car
	toyota = &car{number: 12}
	fmt.Println(toyota)
	var maruti *car
	maruti = new(car)
	(*maruti).number = 14
	fmt.Println((*maruti).number)
	maruti.number = 15

	fmt.Println((*maruti).number)

}

type car struct {
	number int
}
