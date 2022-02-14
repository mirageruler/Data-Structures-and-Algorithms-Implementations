package main

import "fmt"

func main() {
	dataInput := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	fmt.Println("OUTPUT: ", selectionSort(dataInput))
}

func selectionSort(input []int) []int {
	if len(input) == 0 {
		return input
	}

	for i := 0; i < len(input); i++ {
		min := i
		for j := i; j < len(input); j++ {
			if input[j] < input[min] {
				min = j
			}
		}
		input[i], input[min] = input[min], input[i]
	}

	return input
}
