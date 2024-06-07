package recursion

func NestedRecursion(x int) int {
	if x > 100 {
		return x - 10
	}

	return NestedRecursion(NestedRecursion(x + 11))
}
