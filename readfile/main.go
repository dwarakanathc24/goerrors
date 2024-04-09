package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println("the error ", err)
		return
	}
	br, err := io.ReadAll(f)
	fmt.Println("reading a file")

	fmt.Println(string(br))
}
