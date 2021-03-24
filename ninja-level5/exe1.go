package main

import (
	"fmt"
)

type person struct {
	first      string
	last       string
	favflavors []string
}

func main() {

	p1 := person{

		first: "Thanh",
		last:  "Tran",
		favflavors: []string{
			"cafe",
			"money",
			"rum and coke",
		},
	}
	p2 := person{

		first: "Thanh1",
		last:  "Tran2",
		favflavors: []string{
			"Milk",
			"Strawberrry",
			"Juice",
		},
	}
	fmt.Println(p1.first, p1.last)
	for i, v := range p1.favflavors {
		fmt.Println(i, v)

	}
	fmt.Println(p2)

}
