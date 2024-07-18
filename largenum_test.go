package main

import "testing"

func TestFindLargest(t *testing.T) {
	result := FindLargest([]int{7, 12, 3, 45, 23})
	expected := 45
	if result != expected {
		t.Errorf("FindLargest([]int{7, 12, 3, 45, 23}) = %d; want %d", result, expected)
	}

	result = FindLargest([]int{0, -5, 22, 11, -3})
	expected = 22
	if result != expected {
		t.Errorf("FindLargest([]int{0, -5, 22, 11, -3}) = %d; want %d", result, expected)
	}

	result = FindLargest([]int{100, 250, 300, 200, 150})
	expected = 300
	if result != expected {
		t.Errorf("FindLargest([]int{100, 250, 300, 200, 150}) = %d; want %d", result, expected)
	}
}
