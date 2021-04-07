package main

import (
	"fmt"
)

type person struct {
	Frist string
	Last  string
	Age   int
}

func main() {

	p1 := person{
		Frist: "Tran",
		Last:  "Thanh",
		Age:   24,
	}

	p2 := person{
		Frist: "Thanh",
		Last:  "Tran",
		Age:   42,
	}
	people := []person{p1, p2}

	fmt.Println(people)
}
