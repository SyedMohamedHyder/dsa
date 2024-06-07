package recursion

func NestedRecursion(x int) int {
	if x > 100 {
		return x - 10
	} else {
		return NestedRecursion(NestedRecursion(x + 11))
	}
}
