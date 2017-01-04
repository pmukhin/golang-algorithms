package numeric

import "testing"

func TestGreatestCommonDivisor(t *testing.T) {
	if GreatestCommonDivisor(3, 4) != 1 {
		t.Fatal()
	}
	if GreatestCommonDivisor(6, 12) != 6 {
		t.Fatal()
	}
	if GreatestCommonDivisor(2, 4) != 2 {
		t.Fatal()
	}
}
