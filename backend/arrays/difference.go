package arrays

import (
	"slices"
)

func DifferenceUnSorted(arr1, arr2 []int) []int {
	resultArr := make([]int, 0)

	for _, item := range arr1 {
		if !slices.Contains(arr2, item) {
			resultArr = append(resultArr, item)
		}
	}

	return resultArr
}

func DifferenceSorted(arr1, arr2 []int) []int {
	resultArr := make([]int, 0)

	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			resultArr = append(resultArr, arr1[i])
			i++
		} else if arr2[j] < arr1[i] {
			j++
		} else {
			i++
			j++
		}
	}

	resultArr = append(resultArr, arr1[i:]...)

	return resultArr
}
