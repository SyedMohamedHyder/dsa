package recursion

// Time Complexity: O(n)
// Space Complexity: O(n)
func Power(m, n int) int {
	if n > 0 {
		return Power(m, n-1) * m
	}

	return 1
}

// Time Complexity: O(log n)
// Space Complexity: O(1)
func PowerOptimized(m, n int) int {
	if n == 0 {
		return 1
	} else if n%2 == 0 {
		return PowerOptimized(m*m, n/2)
	} else {
		return PowerOptimized(m*m, (n-1)/2) * m
	}
}
