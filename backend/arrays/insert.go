package arrays

func Insert(arr []int, target int) []int {
	tempArr := make([]int, len(arr)+1)
	copy(tempArr, arr)

	i := len(arr) - 1
	for tempArr[i] > target {
		tempArr[i+1] = tempArr[i]
		i--
	}

	tempArr[i+1] = target

	return tempArr
}
