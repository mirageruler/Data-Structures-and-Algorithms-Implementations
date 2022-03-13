package sorting

func Bubble(input []int) []int {
	l := len(input)
	if l == 0 {
		return input
	}

	for l > 1 {
		for j := 0; j < l-1; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
		l--
	}

	return input
}
