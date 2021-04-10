package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {

	u1 := user{
		First: "Tran",
		Age:   24,
	}
	u2 := user{
		First: "Van",
		Age:   25,
	}
	u3 := user{
		First: "Thanh",
		Age:   26,
	}
	users := []user{u1, u2, u3}

	fmt.Println(users)
	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(string(bs))
}
