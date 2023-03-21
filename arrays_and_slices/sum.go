package main

func Sum(numbers [5]int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}
