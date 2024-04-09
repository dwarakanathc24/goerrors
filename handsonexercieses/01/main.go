package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//this is the documentatin

type person struct {
	First string
	Last  string
	Email string
	Loc   []string
}

func main() {
	p := person{
		First: "abhay",
		Last:  "ram",
		Email: "london",
		Loc:   []string{"London", "Ilford", "UK"},
	}

	b, err := json.Marshal(p)
	if err != nil {
		log.Fatalln("JSON is unable to convert into the json----", err)
	}
	fmt.Println(string(b))

}

// addding function calculates the sum of the passed variables
func add(a, b int64) int64 {
	return a + b
}
