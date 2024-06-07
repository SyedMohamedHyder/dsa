package recursion

var a, b int = 1, 1

func Taylor(x, n int) int {
	if n > 0 {
		result := Taylor(x, n-1)
		a *= x
		b *= n
		return result + (a / b)
	}

	return 1
}
