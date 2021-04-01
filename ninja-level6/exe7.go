package main

import (
	"fmt"
)

var x int

func main() {

	a := func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)

		}

	}
	a()
	fmt.Printf("%T\n", a)

	fmt.Println(x)
	fmt.Printf("%T\n", x)

}
