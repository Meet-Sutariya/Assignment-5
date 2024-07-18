package main

import "testing"

// TestFibonacci tests the Fibonacci function with various cases.
func TestFibonacci(t *testing.T) {
	tests := []struct {
		n      int
		expect int
	}{
		{0, 0},            // Base case
		{1, 1},            // Base case
		{2, 1},            // Typical case
		{3, 2},            // Typical case
		{10, 55},          // Typical case
		{20, 6765},        // Typical case
		{50, 12586269025}, // Large number case
	}

	for _, test := range tests {
		result := Fibonacci(test.n)
		if result != test.expect {
			t.Errorf("Fibonacci(%d) = %d; want %d", test.n, result, test.expect)
		}
	}
}
