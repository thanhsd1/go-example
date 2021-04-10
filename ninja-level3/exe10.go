package main

import (
	"fmt"
)

func main() {
	x := [9]int{1, 2, 3, 5, 10}
	fmt.Println(len(x))
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	for i, v := range x {
		fmt.Println(i, v)

	}

}
