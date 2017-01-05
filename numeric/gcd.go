package numeric

// Examples how it works:
// gcd(6, 12) -> (12, 6) -> (6, 0) -> 6
// gcd(3, 4) -> (4, 3) -> (3, 1) -> (1, 0) -> 1
// gcd(2, 4) -> (4, 2) -> (2, 0) -> 2
// gcd(17, 19) -> (19, 17) -> (17, 2) -> (2, 1) -> (1, 0) -> 1

// GreatestCommonDivisor finds Greatest Common Divisor
// of two numbers a and b
func GreatestCommonDivisor(a int, b int) int {
	// b == 0 means a is the GCD
	if 0 == b {
		return a
	}

	// Recursion:
	// In next call a is swapped with b, and b with modulo of a / b
	return GreatestCommonDivisor(b, a%b)
}
