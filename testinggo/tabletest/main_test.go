package main

import "testing"

func TestAdd(t *testing.T) {

	type test struct {
		data   []int
		result int
	}

	tests := []test{
		test{[]int{1, 2}, 3},
		test{[]int{3, 2}, 5},
		test{[]int{3, -2, -1}, 0},
	}

	for _, v := range tests {

	}
	result := add(3, 2)
	expected := 5

	if result != int64(expected) {
		t.Errorf("Add(3, 2) returned %d, expected %d", result, expected)
	}
}

func TestMul(t *testing.T) {
	result := mul(3, 2)
	expected := 6

	if result != int(expected) {
		t.Errorf("mul(3, 2) returned %d, expected %d", result, expected)
	}
}
