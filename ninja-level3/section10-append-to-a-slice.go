package main

import "fmt"

func main() {
	x := []int{1, 3, 4, 6, 8}
	fmt.Println(x)

	x = append(x, 10, 12, 14, 16)
	fmt.Println(x)

	y := []int{20, 22, 24, 26, 28}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[1:]) // delete append
	fmt.Println(x)

	x = append(x[2:]) // delete append
	fmt.Println(x)

	x = append(x[1:], x[2:]...) // delete append
	fmt.Println(x)

}
