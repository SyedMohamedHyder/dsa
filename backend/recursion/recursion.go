package recursion

import "fmt"

func RecurseAndPrint(x int) {
	if x > 0 {
		RecurseAndPrint(x - 1)
		fmt.Print(x)
	}
}

func PrintAndRecurse(x int) {
	if x > 0 {
		fmt.Print(x)
		PrintAndRecurse(x - 1)
	}
}
