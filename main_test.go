package main

import "testing"

func TestFindLargest(t *testing.T) {
	result := FindLargest([]int{1, 2, 3, 4, 5})
	expected := 5
	if result != expected {
		t.Errorf("FindLargest([]int{1, 2, 3, 4, 5}) = %d; want %d", result, expected)
	}

	result = FindLargest([]int{-1, -2, -3, -4, -5})
	expected = -1
	if result != expected {
		t.Errorf("FindLargest([]int{-1, -2, -3, -4, -5}) = %d; want %d", result, expected)
	}
}
