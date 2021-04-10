package main

import (
	"fmt"
)

func main() {
	x := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(x)

	for i, v := range x {
		fmt.Println("This is record for", i)
		for j, k := range v {
			fmt.Println("\t", j, k) // tại sao output của em lại là dòng đầu tiên "bond_jame" vậy anh. Trong Docs thì dòng thứ 3 "no_dr""
		}
	}

}
