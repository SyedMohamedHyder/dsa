package recursion

// Time Complexity: O(n)
// Space Complexity: O(1)
func NCRFormula(n, r int) int {
	return FactorialNLoop(n) / (FactorialNLoop(r) * FactorialNLoop(n-r))
}

// Time Complexity: O(2^n) as it uses tree recursion
// Space Complexity: O(n)
// This uses pascal's triangle formula
func NCRRecurse(n, r int) int {
	if r == 0 || r == n {
		return 1
	}

	return NCRRecurse(n-1, r-1) + NCRRecurse(n-1, r)
}
