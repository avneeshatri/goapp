package main

import (
	"fmt"
	"net/http"
)

func main() {
	//panicDivieBy0()
	userDefinedPanic()
	//httpPanic()
}

func httpPanic() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello Go"))
	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err.Error())
	}
}

func userDefinedPanic() {
	fmt.Println("Start")
	defer fmt.Println("Middle")
	panic("something bad happend")
	fmt.Println("End")
}
func panicDivieBy0() {
	a, b := 1, 0
	ans := a / b
	fmt.Println(ans)
}
