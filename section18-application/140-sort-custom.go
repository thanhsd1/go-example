package main

import (
	"fmt"
)

type person struct {
	first string
	age   int
}

func main() {
	p1 := person{"TranThanh1", 24}
	p2 := person{"TranThanh2", 25}
	p3 := person{"TranThanh3", 26}
	p4 := person{"TranThanh4", 27}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)
}
