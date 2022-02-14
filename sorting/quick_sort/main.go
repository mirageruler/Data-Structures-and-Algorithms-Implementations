package main

import "fmt"

func main() {
	dataInput := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	fmt.Println("OUTPUT: ", quickSort(dataInput))
}

func quickSort(input []int) []int {
	if len(input) == 0 {
		return input
	}

	return input
}
