package sorting

// Bubble sorts an input array
// with bubble sort algorithm.
// This algorithms has constant
// complexity - O(n*n)
func Bubble(input []int) []int {
	length := len(input)
	if length <= 1 {
		return input
	}
	// Iterating through input array
	// in i dimension
	for i := 0; i < length; i++ {
		// Iterating through input array
		// in j dimension
		for j := length - 1; j >= i; j-- {
			// If left element is more than right one
			// then swap them
			if input[i] > input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}

	return input
}
