package sorting

func Merge(input []int) []int {
	return mergeSort(input)
}

func mergeSort(input []int) []int {
	mid := len(input) / 2
	if len(input) < 2 {
		return input
	}

	first := mergeSort(input[:mid])
	second := mergeSort(input[mid:])

	return merge(first, second)
}

func merge(left, right []int) []int {
	result := []int{}
	leftIdx, rightIdx := 0, 0
	for leftIdx < len(left) && rightIdx < len(right) {
		if left[leftIdx] < right[rightIdx] {
			result = append(result, left[leftIdx])
			leftIdx++
		} else if left[leftIdx] > right[rightIdx] {
			result = append(result, right[rightIdx])
			rightIdx++
		}
	}

	for ; leftIdx < len(left); leftIdx++ {
		result = append(result, left[leftIdx])
	}

	for ; rightIdx < len(right); rightIdx++ {
		result = append(result, right[rightIdx])
	}

	return result
}
