package main

import (
	"fmt"
)

func main() {
	f := xyz()
	fmt.Println(f())
}
func xyz() func() int {
	return func() int {

		return 42
	}
}
