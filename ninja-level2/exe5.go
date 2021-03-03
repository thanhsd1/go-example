package main

import (
	"fmt"
)

//raw string literal: Có hai dạng: ký tự chuỗi thô và ký tự chuỗi được giải thích
func main() {
	a := `here is something
	as 
	a 
	raw string
	literal
	"you see"
	another thing`
	fmt.Println(a)
}
