package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type User struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	records, err := readData("/Users/tranthanh/Desktop/filedemo.csv")

	if err != nil {
		fmt.Println(err)
	}

	for _, record := range records {
		user := User{
			firstName: record[0],
			lastName:  record[1],
			age:       int(2),
		}

		fmt.Printf("%s %s  %v\n", user.firstName, user.lastName,
			user.age)
	}
}

func readData(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)

	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
