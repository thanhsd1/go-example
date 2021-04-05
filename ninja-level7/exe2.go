package main

import (
	"fmt"
)

type person struct {
	name string
}

func main() {
	p1 := person{
		name: "Tran Thanh",
	}
	fmt.Println(p1)
	changMe(&p1)
	fmt.Println(p1)

}
func changMe(p *person) {
	p.name = "xxyyzz"
	//(*p).name = "zzyyxx"
	//(*&p).name = "yyzzxx"

}
