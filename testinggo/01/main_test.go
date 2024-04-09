package main

import "testing"

func TestAdd(t *testing.T) {
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
