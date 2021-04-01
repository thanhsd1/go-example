package main

import (
	"fmt"
)

func main() {

	a := func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)

		}

	}
	a()
	fmt.Printf("%T\n", a)
}
