package main

import (
	"fmt"
	"sort"
)

func main() {

	xi := []int{5, 1, 6, 2, 3, 4, 7, 8}
	xs := []string{"My", "Thanh", "Tran", "Name", "Is"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("-------")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs) // Out put sort string nó ra là [Is My Name Thanh Tran] có vấn đề gi v a?


}
