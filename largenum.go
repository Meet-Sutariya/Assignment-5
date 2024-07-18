package main

import (
	"fmt"
)

// MatrixMultiply multiplies two 2x2 matrices.
func MatrixMultiply(a, b [2][2]int) [2][2]int {
	return [2][2]int{
		{a[0][0]*b[0][0] + a[0][1]*b[1][0], a[0][0]*b[0][1] + a[0][1]*b[1][1]},
		{a[1][0]*b[0][0] + a[1][1]*b[1][0], a[1][0]*b[0][1] + a[1][1]*b[1][1]},
	}
}

// MatrixPower raises a 2x2 matrix to the power of n.
func MatrixPower(matrix [2][2]int, n int) [2][2]int {
	result := [2][2]int{{1, 0}, {0, 1}} // Identity matrix
	base := matrix

	for n > 0 {
		if n%2 == 1 {
			result = MatrixMultiply(result, base)
		}
		base = MatrixMultiply(base, base)
		n /= 2
	}

	return result
}

// Fibonacci calculates the nth Fibonacci number using matrix exponentiation.
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	matrix := [2][2]int{{1, 1}, {1, 0}}
	resultMatrix := MatrixPower(matrix, n-1)

	return resultMatrix[0][0]
}

func main() {
	// Example usage of the Fibonacci function.
	fmt.Println("Fibonacci of 50 is:", Fibonacci(50)) // Output: Fibonacci of 50 is: 12586269025
}
