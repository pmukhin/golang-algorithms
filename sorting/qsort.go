package sorting

func QSort(input []int) []int {
	length := len(input)
	if length <= 1 {
		return input
	}
	pivot := input[length-1]

	left := make([]int, 0, length)
	middle := make([]int, 0, length)
	right := make([]int, 0, length)

	for _, v := range input {
		switch {
		case v < pivot:
			left = append(left, v)
		case v == pivot:
			middle = append(middle, v)
		case v > pivot:
			right = append(right, v)
		}
	}

	left = QSort(left)
	right = QSort(right)

	left = append(left, middle...)

	return append(left, right...)
}
