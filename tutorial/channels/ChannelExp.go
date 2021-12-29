package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)

	//	for i := 0; i < 5; i++ {
	wg.Add(2)

	go func(ch <-chan int) {
	f1:
		for {

			if i, ok := <-ch; ok {
				fmt.Println("Read from channel:", i)
			} else {
				fmt.Println("Channel is closed")
				break f1
			}
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {

		ch <- 43
		ch <- 37
		fmt.Println("Pushed to channel:")
		close(ch)
		wg.Done()
	}(ch)
	//}
	wg.Wait()
}
