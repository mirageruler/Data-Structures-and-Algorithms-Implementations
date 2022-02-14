package main

import (
	"fmt"
)

func main() {
	dataInput := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	fmt.Println("OUTPUT: ", bubbleSort(dataInput))
}

func bubbleSort(input []int) []int {
	if len(input) == 0 {
		return input
	}

	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input)-1-i; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	return input
}
