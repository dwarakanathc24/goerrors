package main

import "fmt"

func add(a, b int64) int64 {
	return a + b
}
func mul(t ...int) int {
	s := 1
	for _, i := range t {
		s = s * i
	}
	return s
}
func main() {
	fmt.Println(add(3, 2))
	fmt.Println(mul(3, 2))

}
