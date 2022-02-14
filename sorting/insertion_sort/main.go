package main

import "fmt"

func main() {
	dataInput := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	fmt.Println("OUTPUT: ", insertionSort(dataInput))
}

func insertionSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}

	for i := 1; i < len(input); i++ {
		for j := i; j > 0 && input[j-1] > input[j]; j-- {
			input[j-1], input[j] = input[j], input[j-1]
		}
	}

	return input
}
