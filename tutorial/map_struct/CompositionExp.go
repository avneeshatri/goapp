package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `required max:100 dummy`
	Origin string
}

type Bird struct {
	Animal
	CanFly bool
}

func main() {
	sparrow := Bird{Animal: Animal{Name: "Sparrow", Origin: "India"},
		CanFly: true,
	}

	//sparrow.Name = "Sparrow"
	//sparrow.Origin = "India"

	fmt.Println(sparrow)

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
