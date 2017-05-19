package binary

import "testing"

func TestTreeNode_Append(t *testing.T) {
	tn := New(50)
	values := []int{17, 72, 12, 23, 54, 76, 9, 14, 19, 67}

	for _, value := range values {
		tn.Append(New(value))
	}

	{
		var _values = make([]int, 0, 11)
		IterLeft(tn, &_values)

		if len(_values) != 11 {
			t.Error("To few values were found")
		}

		// as we iter from left to right
		for i, v := range _values {
			if i == 0 || i == (len(_values)-1) {
				continue
			}

			if v < _values[i-1] {
				t.Error("%d goes after %d", v, _values[i])
			}
		}
	}

	{
		var _values = make([]int, 0, 11)
		IterRight(tn, &_values)

		if len(_values) != 11 {
			t.Error("To few values were found")
		}

		// as we iter from right to left
		for i, v := range _values {
			if i == 0 || i == (len(_values)-1) {
				continue
			}

			if v > _values[i-1] {
				t.Error("%d goes after %d", v, _values[i])
			}
		}
	}

	{
		var _values = make([]int, 0, 11)
		Iter(tn, &_values)

		if len(_values) != 11 {
			t.Error("To few values were found")
		}

		var mustBeValues = []int{50, 17, 12, 9, 14, 23, 19, 72, 54, 67, 76}
		for i, v := range _values {
			if mustBeValues[i] != v {
				t.Error("Wrong order of values")
			}
		}
	}

}
