package arrays

import (
	"slices"
)

// Time Complexity : O(n^2)
// Space Complexity : O(m+n)
func IntersectionUnsorted(arr1, arr2 []int) []int {
	resultArr := make([]int, 0)

	for _, item := range arr1 {
		if slices.Contains(arr2, item) {
			resultArr = append(resultArr, item)
		}
	}

	return resultArr
}

// Time Complexity : O(m+n)
// Space Complexity : O(m+n)
func IntersectionSorted(arr1, arr2 []int) []int {
	resultArr := make([]int, 0)

	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			i++
		} else if arr2[j] < arr1[i] {
			j++
		} else {
			resultArr = append(resultArr, arr1[i])
			i++
			j++
		}
	}

	return resultArr
}
