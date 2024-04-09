package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//writing a file

	f, err := os.Create("courses.txt")

	if err != nil {
		fmt.Println(err)
		return

	}

	defer f.Close()

	s := strings.NewReader(" Java GO Python Ruby")
	io.Copy(f, s)

	//read a file

	f, err = os.Open("courses.txt")

	if err != nil {
		fmt.Println("error happend ", err)
	}

	br, err := io.ReadAll(f)
	fmt.Println(string(br))
}
