package main

import (
	"encoding/csv"
	"fmt"

	"os"
)

func main() {

	records := [][]string{
		{"ho", "ten", "nghe nghiep"},
		{"Tran", "Thanh", "qc"},
		{"Tran1", "Thanh1", "qc1"},
		{"Tran12", "Thanh12", "qc12"},
	}

	f, err := os.Create("object.csv")
	defer f.Close()

	if err != nil {

		fmt.Println("file die", err)
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(records)

	if err != nil {
		fmt.Println(err)
	}
}
