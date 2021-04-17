package main

import (
	"fmt"
)

type person struct {
	first string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Prinln("Huhu")
}

func saySomething(h human)

func main() {

	p1 := person{
		first: "Thanh",
	}
	saySomething(&p1)

	p1.speak()

}
