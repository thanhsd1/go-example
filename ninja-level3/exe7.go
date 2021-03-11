package main

import (
	"fmt"
)

func main() {
	x := "tran van thanh"
	if x == "van" {
		fmt.Println(x)
	} else if x == "tran van thanh" {
		fmt.Println("tran van thanh", x)
	} else {
		fmt.Println("xyz")
	}

}
