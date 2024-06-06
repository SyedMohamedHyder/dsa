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

	fmt.Println("****Sum with Global Var****")
	fmt.Println(recursion.SumWithGlobalVar(5))

	fmt.Println("****Tail Recurse and Print****")
	recursion.TailRecurseAndPrint(5)

	fmt.Println("****Loop and Print****")
	recursion.LoopAndPrint(5)
}
