package sorting

import "testing"

func TestQSort(t *testing.T) {
	input := []int{3, 2, 1}
	output := QSort(input)

	for k, v := range output {
		if v != k+1 {
			t.Fatal("Must equal")
		}
	}
}
