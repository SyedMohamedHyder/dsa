package recursion

import "fmt"

func TailRecurseAndPrint(x int) {
	if x > 0 {
		fmt.Println(x)
		TailRecurseAndPrint(x - 1)
	}
}

func LoopAndPrint(x int) {
	for x > 0 {
		fmt.Println(x)
		x--
	}
}
