package main

import (
	"fmt"
)

func main() {
	id := 1997
	for {

		if id > 2021 {
			break // vòng lặp dừng nếu id > 2021
		}
		fmt.Println(id)
		id++
	}

}
