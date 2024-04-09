package main

import "fmt"

func main() {
	fmt.Println(" the f calling")
	f()
}

func f() {
	defer func() {

		if r := recover(); r != nil {
			fmt.Println("recovering from panic:", r)
		}
	}()

	g(1)
}

func g(i int32) {
	if i > 3 {
		fmt.Println(" panicking ")
		panic(fmt.Sprintf("%v", i))
	}

	// defer fmt.Println(" defering i:", i)
	fmt.Println(" the value of i:", i)
	g(i + 1)
}
