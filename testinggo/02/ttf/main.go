package main

func add(a ...int) int {
	sum := 0
	for _, i := range a {
		sum = +i
	}
	return sum
}
