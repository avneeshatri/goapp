package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("start prog")
	defer fmt.Println("middle prog")
	defer fmt.Println("end prog")
	deferResource()
	deferValueType()
	deferRefType()
}

func deferRefType() {
	m := map[string]int{"delhi": 1, "meerut": 2}
	defer fmt.Println(m)
	m["delhi"] = 3
}

func deferValueType() {
	a := "start"
	defer fmt.Println(a)
	a = "end"
}

func deferResource() {
	res, err := http.Get("https://www.google.com/robots.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", robots)
}
