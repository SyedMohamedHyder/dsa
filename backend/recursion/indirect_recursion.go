// Indirect recursion is when a func A calls func B which again calls func A.
package recursion

import "fmt"

func IndirectRecurseA(x int) {
	if x > 0 {
		fmt.Printf("%d ", x)
		IndirectRecurseB(x - 1)
	}
}

func IndirectRecurseB(x int) {
	if x > 0 {
		fmt.Printf("%d ", x)
		IndirectRecurseA(x / 2)
	}
}
