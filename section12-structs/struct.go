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

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

}
