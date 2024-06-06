// Tree recursion is when a function call occurs more than once.
// The O(n) in this case is 2^n.
package recursion

import "fmt"

func TreeRecurseAndPrint(x int) {
	if x > 0 {
		fmt.Print(x)
		TreeRecurseAndPrint(x - 1)
		TreeRecurseAndPrint(x - 1)
	}
}
