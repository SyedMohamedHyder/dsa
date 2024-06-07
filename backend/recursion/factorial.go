package recursion

// Time Complexity: O(n)
// Space Complexity: O(n)
func FactorialNRecurse(n int) int {
	if n > 1 {
		return FactorialNRecurse(n-1) * n
	}

	return 1
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func FactorialNLoop(n int) int {
	result := 1

	for i := 1; i <= n; i++ {
		result *= i
	}

	return result
}
