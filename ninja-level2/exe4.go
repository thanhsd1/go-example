package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	//bai tap nay em chua hieu a
	b := a << 1
	fmt.Printf("%d\t%b\t%#x", b, b, b)
}
