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

	fmt.Println("****TailLoop and Print****")
	recursion.LoopAndPrint(5)

	fmt.Println("****Head Recurse and Print****")
	recursion.HeadRecurseAndPrint(5)

	fmt.Println("****Head Loop and Print****")
	recursion.HeadLoopAndPrint(5)

	fmt.Println("****Tree Recurse and Print****")
	recursion.TreeRecurseAndPrint(5)
}
