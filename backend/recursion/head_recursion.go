// Head Recursion is when the function calls itself at the beginning
package recursion

import "fmt"

func HeadRecurseAndPrint(x int) {
	if x > 0 {
		HeadRecurseAndPrint(x - 1)
		fmt.Print(x)
	}
}

func HeadLoopAndPrint(x int) {
	i := 1
	for i < x {
		fmt.Print(x)
		i++
	}
}
