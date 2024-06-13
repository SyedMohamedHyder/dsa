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

	arr = []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	fmt.Println("****Is Sorted (True)****")
	fmt.Println(arrays.IsSorted(arr))

	arr = []int{19, 17, 15, 13, 11, 9, 7, 5, 3, 1}
	fmt.Println("****Is Sorted (False)****")
	fmt.Println(arrays.IsSorted(arr))

	arr = []int{1, -1, 2, -2, 3, -3, 4, -4, 5, -5, -6}
	fmt.Println("****Move Negatives****")
	arrays.ArrangeNegatives(arr)
	fmt.Println(arr)

	arr1 := []int{1, 9, 11, 13, 15, 17, 19}
	arr2 := []int{2, 4, 6, 8, 10, 16, 18, 20}
	fmt.Println("****Merge Arrays****")
	resultArr = arrays.MergeArrays(arr1, arr2)
	fmt.Println(resultArr)

	arr1 = []int{1, 2, 11, 6, 15, 7, 19}
	arr2 = []int{1, 9, 11, 13, 15, 17, 19}
	fmt.Println("****Union UnSorted****")
	resultArr = arrays.UnionUnSorted(arr1, arr2)
	fmt.Println(resultArr)

	arr1 = []int{1, 2, 6, 7, 11, 15, 19}
	arr2 = []int{1, 9, 11, 13, 15, 17, 19, 20}
	fmt.Println("****Union Sorted****")
	resultArr = arrays.UnionSorted(arr1, arr2)
	fmt.Println(resultArr)

	arr1 = []int{19, 2, 11, 6, 15, 7, 1, 20}
	arr2 = []int{1, 9, 11, 13, 15, 17, 19}
	fmt.Println("****Intersection UnSorted****")
	resultArr = arrays.IntersectionUnsorted(arr1, arr2)
	fmt.Println(resultArr)

	arr1 = []int{1, 2, 6, 7, 11, 15, 19, 20}
	arr2 = []int{1, 9, 11, 13, 15, 17, 19}
	fmt.Println("****Intersection Sorted****")
	resultArr = arrays.IntersectionSorted(arr1, arr2)
	fmt.Println(resultArr)

	arr1 = []int{19, 2, 11, 6, 15, 7, 1}
	arr2 = []int{1, 9, 11, 13, 15, 17, 19, 20}
	fmt.Println("****Difference UnSorted****")
	resultArr = arrays.DifferenceUnSorted(arr1, arr2)
	fmt.Println(resultArr)

	arr1 = []int{1, 2, 6, 7, 11, 15, 19}
	arr2 = []int{1, 9, 11, 13, 15, 17, 19, 20}
	fmt.Println("****Difference Sorted****")
	resultArr = arrays.DifferenceSorted(arr1, arr2)
	fmt.Println(resultArr)
}
