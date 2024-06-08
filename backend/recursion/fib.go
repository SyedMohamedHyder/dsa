package recursion

// Time Complexity: O(2^n) because tree recursion is used
// Space Complexity: O(n)
func FibRecurse(n int) int {
	if n > 1 {
		return FibRecurse(n-1) + FibRecurse(n-2)
	}

	return n
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func FibLoop(n int) int {
	a, b := 0, 1

	var result int
	for i := 2; i <= n; i++ {
		result = a + b
		a = b
		b = result
	}

	return result
}

var fib map[int]int = map[int]int{0: 0, 1: 1}

// Time Complexity: O(n)
// Space Complexity: O(n)
func FibMemoization(n int) int {
	if val, ok := fib[n]; ok {
		return val
	}

	fib[n-1] = FibMemoization(n - 1)
	fib[n-2] = FibMemoization(n - 2)
	fib[n] = fib[n-1] + fib[n-2]

	return fib[n]
}
