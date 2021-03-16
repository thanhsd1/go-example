package main

import "fmt"

func main() {
	x := map[string]int{ // Cách để khai báo kiểu chữ và kiểu số trong hàm map

		"Van Thanh":  97,
		"Tran Thanh": 79,
	}
	fmt.Println(x)

	delete(x, "Van Thanh")
	fmt.Println(x)

	delete(x, "Tran Thanh")
	fmt.Println(x)
} // output sẽ là map[]
