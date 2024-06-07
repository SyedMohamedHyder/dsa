package recursion

// Time Complexity: O(n)
// Space Complexity: O(n)
func SumNRecurse(n int) int {
	if n > 0 {
		return SumNRecurse(n-1) + n
	}

	return 0
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func SumNLoop(n int) int {
	result := 0

	for i := 1; i <= n; i++ {
		result += i
	}

	return result
}

// Time Complexity: O(1)
// Space Complexity: O(1)
func SumNFormula(n int) int {
	return n * (n + 1) / 2
}
