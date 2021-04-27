package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	data := []byte("golang vo dich")
	err := ioutil.WriteFile("thanh1.txt", data, 0777)
	if err != nil {

		fmt.Println(err)
	}

	data, err = ioutil.ReadFile("thanh1.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
