package recursion

var globalVar int = 0

func SumWithGlobalVar(n int) int {
	if n > 0 {
		globalVar++
		return SumWithGlobalVar(n-1) + globalVar
	}

	return 0
}
