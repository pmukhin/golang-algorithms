package sorting

import "testing"

func TestBubble(t *testing.T) {
	input := []int{3, 2, 1}
	output := Bubble(input)

	for k, v := range output {
		if v != k+1 {
			t.Fatal("Must equal")
		}
	}
}
