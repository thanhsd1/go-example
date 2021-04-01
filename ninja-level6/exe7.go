package main

import (
	"fmt"
)

var x int
var g func()

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

	g = a
	g()

	fmt.Printf("Day la g %T\n", g)
	fmt.Println("Done")

}
