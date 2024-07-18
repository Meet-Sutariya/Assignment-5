package main

func FindLargest(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	largest := numbers[0]
	for _, num := range numbers {
		if num > largest {
			largest = num
		}
	}
	return largest
}
