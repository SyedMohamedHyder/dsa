package recursion

var a, b float64 = 1, 1

func Taylor(x, n float64) float64 {
	if n > 0 {
		result := Taylor(x, n-1)
		a *= x
		b *= n
		return result + (a / b)
	}

	return 1
}
