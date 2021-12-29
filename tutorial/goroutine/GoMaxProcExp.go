package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(100)
	fmt.Printf("%v\n", runtime.GOMAXPROCS(-1))
}
