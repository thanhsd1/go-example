package main

import (
	"fmt"
)

func main() { // khai báo kiểu string
	x := "van thanh 3"
	switch x {
	case "van thanh":
		fmt.Println("tran van thanh")
	case "van thanh 1":
		fmt.Println("tran van thanh 1")
	case "van thanh 2":
		fmt.Println("tran van thanh 2")
	case "van thanh 3":
		fmt.Println("tran van thanh 3")
	case "van thanh 4":
		fmt.Println("tran van thanh 4")

	}
}
