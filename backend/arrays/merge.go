package arrays

// Time Complexity : O(m+n)
// Space Complexity : O(m+n)
func MergeArrays(arr1, arr2 []int) []int {
	resultArr := make([]int, len(arr1)+len(arr2))

	i, j, k := 0, 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			resultArr[k] = arr1[i]
			i++
		} else {
			resultArr[k] = arr2[j]
			j++
		}

		k++
	}

	for ; i < len(arr1); i++ {
		resultArr[k] = arr1[i]
		k++
	}

	for ; j < len(arr2); j++ {
		resultArr[k] = arr2[j]
		k++
	}

	return resultArr
}
