package main

import (
	"fmt"

	"github.com/SyedMohamedHyder/dsa/backend/recursion"
)

func main() {
	fmt.Println("****Recurse and Print****")
	recursion.RecurseAndPrint(5)
	fmt.Println("****Print and Recurse****")
	recursion.PrintAndRecurse(5)
}
