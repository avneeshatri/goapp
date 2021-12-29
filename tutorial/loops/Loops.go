package main

import (
	"fmt"
)

func main() {
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}

	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	k := 0
	for {
		fmt.Println(k)
		k++

		if k > 3 {
			break
		}
	}

	s := []int{1, 2, 3}

	for k, v := range s {
		fmt.Println(k, v)
	}

	m := map[string]int{"delhi": 1, "meerut": 2}

	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("Loop break")
Loop:
	for i := 0; i < 3; i = i + 1 {
		for j := 0; j < 3; j = j + 1 {
			fmt.Println(i * j)
			if i*j > 3 {
				break Loop
			}
		}
	}
}
