package main

import (
	"fmt"

	"github.com/SyedMohamedHyder/dsa/backend/arrays"
)

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	fmt.Println("****Linear Search****")
	fmt.Println(arrays.LinearSearch(arr, 13))
	fmt.Println(arr)

	fmt.Println("****Transposition Linear Search****")
	fmt.Println(arrays.TranspositionLinearSearch(arr, 13))
	fmt.Println(arr)

	fmt.Println("****Move In Front Linear Search****")
	fmt.Println(arrays.MoveInFrontLinearSearch(arr, 13))
	fmt.Println(arr)

	arr = []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	fmt.Println("****Binary Search Loop****")
	fmt.Println(arrays.BinarySearchLoop(arr, 19))
	fmt.Println(arr)

	fmt.Println("****Binary Search Recurse****")
	fmt.Println(arrays.BinarySearchRecurse(arr, 19, 0, len(arr)-1))
	fmt.Println(arr)

	arr = []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	fmt.Println("****Reverse Copy Back****")
	arrays.ReverseCopyBack(arr)
	fmt.Println(arr)

	fmt.Println("****Reverse Swap****")
	arrays.ReverseSwap(arr)
	fmt.Println(arr)

	arr = []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	fmt.Println("****Left Shift Arr****")
	arrays.LeftShiftArr(arr, 2)
	fmt.Println(arr)

	arr = []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	fmt.Println("****Right Shift Arr****")
	arrays.RightShiftArr(arr, 2)
	fmt.Println(arr)

	arr = []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	fmt.Println("****Rotate Left****")
	arrays.RotateLeft(arr, 2)
	fmt.Println(arr)

	arr = []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	fmt.Println("****Rotate Right****")
	arrays.RotateRight(arr, 2)
	fmt.Println(arr)

	arr = []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	fmt.Println("****Insert****")
	resultArr := arrays.Insert(arr, 12)
	fmt.Println(resultArr)
}
