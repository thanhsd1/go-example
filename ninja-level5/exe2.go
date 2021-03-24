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

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favflavors {
			fmt.Println(i, val)
		}
		fmt.Println("------")

	}

}

/* Output
Tran
Thanh
Tran
0 cafe
1 money
2 rum and coke
------
Tran1
Thanh1
Tran1
0 Milk
1 Strawberrry
2 Juice
*/
