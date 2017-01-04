package numeric

import "testing"

func TestLeastCommonMultiple(t *testing.T) {
	cases := make(map[int][]int)

	cases[12] = []int{3, 4}
	cases[3] = []int{1, 3}
	cases[10] = []int{2, 5}

	for response, args := range cases {
		if LeastCommonMultiple(args[0], args[1]) != response {
			t.Fatal()
		}
	}
}
