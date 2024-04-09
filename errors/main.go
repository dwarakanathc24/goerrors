package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println(" The error case..!")

	i, j := fmt.Println("hello")
	fmt.Println(i, j)

	//taking input from keyboard

	/* var a, b, c string
	fmt.Println("A:")

	_, err := fmt.Scan(&a)
	if err != nil {
		panic(err)
	}

	fmt.Println("B:")
	_, err = fmt.Scan(&b)
	if err != nil {
		panic(err)
	}

	fmt.Println("C:")
	_, err = fmt.Scan(&c)
	if err != nil {
		panic(err)
	}

	fmt.Println(a, b, c) */

	//creating a file

	f, err := os.Create("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer f.Close()
	r := strings.NewReader("my name is abhay ram")

	b, err := io.Copy(f, r)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
