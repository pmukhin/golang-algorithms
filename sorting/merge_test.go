package sorting

import (
	"math/rand"
	"testing"
)

func TestMergeSort(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var j int

	for i := range input {
		j = rand.Intn(len(input) - 1)
		input[i], input[j] = input[j], input[i]
	}

	output := MergeSort(input)

	for i := 0; i < 10; i++ {
		if output[i] != i {
			t.Fatalf("output[%d] must be equal to %d", i, i)
		}
	}
}
