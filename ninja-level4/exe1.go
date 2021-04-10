package main

import (
	"fmt"
)

func main() {
	x := [5]int{1, 3, 5, 6, 10}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println(x)
	fmt.Printf("%T\n", x)

}
