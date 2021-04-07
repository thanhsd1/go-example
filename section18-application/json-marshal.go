package main

import (
	"encoding/json"
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

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)

	}

	fmt.Println(string(bs))
}
