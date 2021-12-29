package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Start")
	panicer()
	fmt.Println("End")
}

func panicer() {
	fmt.Println("panic will happen")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error :", err)
			panic(err)
		}
	}()

	panic("something bad happend")

	fmt.Println("panic happened")
}
