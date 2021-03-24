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
		last:  "Tran1",
		favflavors: []string{
			"Milk",
			"Strawberrry",
			"Juice",
		},
	}
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}

}
