package sorting

// Merge merges two slices
func Merge(left, right []int) []int {
	// First creating a result array with summed lengths
	// of both slices
	result := make([]int, 0, len(left)+len(right))

	// Iterating until both arrays are empty
	for len(left) > 0 || len(right) > 0 {
		// If left is empty - append all elements
		// from right to result
		if len(left) == 0 {
			result = append(result, right...)
			// It's done now
			break
		}

		// If right is empty we append left to result
		if len(right) == 0 {
			result = append(result, left...)
			// It's done here
			break
		}

		// If left[0] is less or equals it goes to result
		if left[0] <= right[0] {
			result = append(result, left[0])
			// And we remove first element
			left = left[1:]
		} else {
			// Same is here.
			result = append(result, right[0])
			right = right[1:]
		}
	}

	// Return result
	return result
}

// MergeSort sorts array with complexity
// O (n log n)
func MergeSort(input []int) []int {
	// First let's get length
	length := len(input)

	// If it's less or equals one return input
	if length <= 1 {
		return input
	}

	// Getting middle element index
	middle := length / 2

	// Sorting all elements from left side
	left := MergeSort(input[middle:])
	// Sorting all elements from right side
	right := MergeSort(input[:middle])

	// Merge and return
	return Merge(left, right)
}
