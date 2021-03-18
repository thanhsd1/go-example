package main

import (
	"fmt"
)

func main() {
	x1 := []string{"James", "Bond", "Shaken, not stirred"}
	x2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	fmt.Println(x1)
	fmt.Println(x2)

	x12 := [][]string{x1, x2}
	fmt.Println(x12)

	for i, v := range x12 {
		fmt.Println("record: ", i)
		for j, val := range v {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}

}
