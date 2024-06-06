// Function that calls itself at the end is tail recursion
package recursion

import "fmt"

func TailRecurseAndPrint(x int) {
	if x > 0 {
		fmt.Print(x)
		TailRecurseAndPrint(x - 1)
	}
}

func LoopAndPrint(x int) {
	for x > 0 {
		fmt.Print(x)
		x--
	}
}
