package arrays

func LinearSearch(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}

	// if target is not found
	return -1
}

func TranspositionLinearSearch(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			if i != 0 {
				arr[i-1], arr[i] = arr[i], arr[i-1]
			}
			return i
		}
	}

	return -1
}

func MoveInFrontLinearSearch(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			arr[i], arr[0] = arr[0], arr[i]
			return i
		}
	}

	return -1
}
