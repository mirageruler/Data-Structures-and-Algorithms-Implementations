package sorting

func Selection(input []int) []int {
	l := len(input)
	if l == 1 {
		return input
	}

	for i := 0; i < l; i++ {
		minIndex := i // assume the index of the smallest element is at the current first index.
		for j := i + 1; j < l; j++ {
			if input[j] < input[minIndex] {
				// re-assign the index of min so we can have the min index to perform the next compare
				// and to do the swap after this inner loop
				minIndex = j
			}
		}
		// swap whatever value at the min index with the current first index of the array
		// at each iteration of the outter loop, the current first index will be closer to the end by 1 unit.
		input[i], input[minIndex] = input[minIndex], input[i]
	}

	return input
}
