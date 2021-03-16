package main

import "fmt"

func main() {
	x := map[string]int{ // Cách để khai báo kiểu chữ và kiểu số trong hàm map

		"Van Thanh":  97,
		"Tran Thanh": 79,
	}
	fmt.Println(x)

	fmt.Println(x["Van Thanh"])

	fmt.Println(x["Van Thanh1"]) // Giá trị này sẽ = 0 vì chưa được khai báo

	fmt.Println(x["Tran Thanh"])

}
