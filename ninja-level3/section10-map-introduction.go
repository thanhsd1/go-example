package main

import (
	"fmt"
)

func main() {
	x := map[string]int{ // Cách để khai báo kiểu chữ và kiểu số trong hàm map

		"Van Thanh":  97,
		"Tran Thanh": 79,
	}
	fmt.Println(x)

	fmt.Println(x["Van Thanh"])

	fmt.Println(x["Van Thanh1"]) // Giá trị này sẽ = 0 vì chưa được khai báo

	fmt.Println(x["Tran Thanh"])

	v, ok := x["Van Thanh1"]
	fmt.Println(v)
	fmt.Println(ok)

	x["Van Thanh2"] = 44

	if v, ok := x["Van Thanh1"]; ok {
		fmt.Println(v, ok)

	}

	for k, v := range x {
		fmt.Println(k, v)
	}

	x1 := []int{1, 3, 5, 7, 10}
	fmt.Println(x1)

	for i, v := range x1 {
		fmt.Println(i, v)
	}
}
