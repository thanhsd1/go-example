package main

import "fmt"

func main() {
	x := []int{5, 4, 3, 1, 8}
	fmt.Println(x[1]) // lấy ra 1 phần tử trong mảng
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:5]) //  ở đây có nghĩa là phần tử sau trừ cho phần tử trước. Đúng không vậy anh??

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
}
