package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
		fmt.Println("anonymous func called")
	}()

	fmt.Println("Start of main")
	doSomething()
	fmt.Println("End of main")
}

func doSomething() {
	fmt.Println("Start of doSomething")
	m := map[string]string{"name": "ramana", "age": "52"}
	// Triggering a panic
	panic(m)
	fmt.Println("End of doSomething") // This won't be executed
}
