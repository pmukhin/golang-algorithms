package sorting

import "math/rand"

// QSort sorts an input array
// Best time for this algorithm is O(n log n)
func QSort(input []int) []int {
	length := len(input)
	// We will use a tail call here
	// so we need a point to stop
	if length <= 1 {
		return input
	}
	// First let's find a pivot element
	pivot := input[rand.Intn(length - 1)]

	// Now let's create 3 arrays with input array length
	left := make([]int, 0, length)
	middle := make([]int, 0, length)
	right := make([]int, 0, length)

	// Iterating through input
	for _, v := range input {
		switch {
		// If value is more than pivot
		// it goes to the left
		case v < pivot:
			left = append(left, v)
		// If it equals to pivot
		// it goes to the middle
		case v == pivot:
			middle = append(middle, v)
		// If it is more than pivot
		// it goes to the right
		case v > pivot:
			right = append(right, v)
		}
	}
	// Recursively sort resulting arrays
	// We don't need to sort the middle,
	// as all values in it are equal
	left = QSort(left)
	right = QSort(right)

	// Append both middle and right to left
	left = append(left, middle...)
	// And return
	return append(left, right...)
}
