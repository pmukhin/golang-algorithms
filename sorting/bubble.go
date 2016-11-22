package sorting

func Bubble(input []int) []int {
	length := len(input)
	if length <= 1 {
		return input
	}
	for i := 0; i < length; i++ {
		for j := length - 1; j >= i; j-- {
			if input[i] > input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}

	return input
}
