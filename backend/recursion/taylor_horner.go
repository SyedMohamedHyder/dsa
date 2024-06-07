package recursion

var result float64 = 1

func TaylorRecurse(x, n float64) float64 {
	if n > 0 {
		result = 1 + (x * result / n)
		return TaylorRecurse(x, n-1)
	}

	return result
}

func TaylorLoop(x, n float64) float64 {
	result = 1

	for n > 0 {
		result = 1 + (x * result / n)
		n--
	}

	return result
}
