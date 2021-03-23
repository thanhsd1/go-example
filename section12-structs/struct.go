package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{

		first: "Van",
		last:  "Thanh",
		age:   27,
	}

	p2 := person{

		first: "Van1",
		last:  "Thanh1",
		age:   24,
	}
	fmt.Println(p1)
	fmt.Println(p2)

}
