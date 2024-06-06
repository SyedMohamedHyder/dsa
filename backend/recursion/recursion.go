package recursion

import "fmt"

func RecurseAndPrint(x int) {
	if x > 0 {
		RecurseAndPrint(x - 1)
		fmt.Println(x)
	}
}

func PrintAndRecurse(x int) {
	if x > 0 {
		fmt.Println(x)
		PrintAndRecurse(x - 1)
	}
}
