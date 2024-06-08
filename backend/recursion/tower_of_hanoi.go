package recursion

// Time Complexity: O(2^n) as it uses tree recursion
// Space Complexity: O(n)
func TOH(n, A, B, C int) [][]int {

	result := make([][]int, 0)
	if n > 0 {
		result = append(result, TOH(n-1, A, C, B)...)
		result = append(result, []int{A, C})
		result = append(result, TOH(n-1, B, A, C)...)
	}

	return result
}
