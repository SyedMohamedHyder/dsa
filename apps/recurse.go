package main

import (
	"fmt"

	"github.com/SyedMohamedHyder/dsa/backend/recursion"
)

func main() {
	fmt.Println("****Recurse and Print****")
	recursion.RecurseAndPrint(5)
	fmt.Println()

	fmt.Println("****Print and Recurse****")
	recursion.PrintAndRecurse(5)
	fmt.Println()

	fmt.Println("****Sum with Global Var****")
	fmt.Println(recursion.SumWithGlobalVar(5))

	fmt.Println("****Tail Recurse and Print****")
	recursion.TailRecurseAndPrint(5)
	fmt.Println()
	fmt.Println("****TailLoop and Print****")
	recursion.LoopAndPrint(5)
	fmt.Println()

	fmt.Println("****Head Recurse and Print****")
	recursion.HeadRecurseAndPrint(5)
	fmt.Println()

	fmt.Println("****Head Loop and Print****")
	recursion.HeadLoopAndPrint(5)
	fmt.Println()

	fmt.Println("****Tree Recurse and Print****")
	recursion.TreeRecurseAndPrint(5)
	fmt.Println()

	fmt.Println("****Indirect Recurse****")
	recursion.IndirectRecurseA(20)
	fmt.Println()

	fmt.Println("****Nested Recurse****")
	fmt.Println(recursion.NestedRecursion(95))

	fmt.Println("****Sum N Recurse****")
	fmt.Println(recursion.SumNRecurse(5))

	fmt.Println("****Sum N Loop****")
	fmt.Println(recursion.SumNLoop(5))

	fmt.Println("****Sum N Formula****")
	fmt.Println(recursion.SumNFormula(5))

	fmt.Println("****Factorial N Recurse****")
	fmt.Println(recursion.FactorialNRecurse(5))

	fmt.Println("****Factorial N Loop****")
	fmt.Println(recursion.FactorialNLoop(5))

	fmt.Println("****Power N Recurse****")
	fmt.Println(recursion.Power(2, 5))

	fmt.Println("****Power N Optimized****")
	fmt.Println(recursion.Power(2, 5))

	fmt.Println("****Taylor****")
	fmt.Println(recursion.Taylor(4, 10))

	fmt.Println("****Taylor Recurse****")
	fmt.Println(recursion.TaylorRecurse(4, 10))

	fmt.Println("****Taylor Loop****")
	fmt.Println(recursion.TaylorLoop(4, 10))

	fmt.Println("****Fib Recurse****")
	fmt.Println(recursion.FibRecurse(10))

	fmt.Println("****Fib Loop****")
	fmt.Println(recursion.FibLoop(10))

	fmt.Println("****Fib Memoization****")
	fmt.Println(recursion.FibMemoization(10))

	fmt.Println("****nCr formula****")
	fmt.Println(recursion.NCRFormula(4, 2))

	fmt.Println("****nCr recurse****")
	fmt.Println(recursion.NCRRecurse(4, 2))
}
