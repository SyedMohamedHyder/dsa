package arrays

import (
	"slices"
)

// Time Complexity : O(n^2)
// Space Complexity : O(m+n)
func UnionUnSorted(arr1 []int, arr2 []int) []int {
	resultArr := make([]int, len(arr1))

	copy(resultArr, arr1)

	for _, item := range arr2 {
		if !slices.Contains(resultArr, item) {
			resultArr = append(resultArr, item)
		}
	}

	return resultArr
}

// Time Complexity : O(m+n)
// Space Complexity : O(m+n)
func UnionSorted(arr1, arr2 []int) []int {
	resultArr := make([]int, 0)
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			resultArr = append(resultArr, arr1[i])
			i++
		} else if arr2[j] < arr1[i] {
			resultArr = append(resultArr, arr2[j])
			j++
		} else {
			resultArr = append(resultArr, arr1[i])
			i++
			j++
		}
	}

	resultArr = append(resultArr, arr1[i:]...)
	resultArr = append(resultArr, arr2[j:]...)

	return resultArr
}
