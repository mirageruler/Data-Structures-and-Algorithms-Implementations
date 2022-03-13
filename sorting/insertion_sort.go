package sorting

func Insertion(input []int) []int {
	l := len(input)
	if l == 0 {
		return input
	}

	for i := 0; i < l; i++ {
		temp := input[i]
		j := i
		for ; j > 0 && temp < input[j-1]; j-- {
			input[j] = input[j-1]
		}
		input[j] = temp
	}
	return input
}
