package arrays

func RotateLeft(arr []int, d int) {
	d %= len(arr)

	tempArr := make([]int, len(arr))
	copy(tempArr, arr[d:])
	copy(tempArr[len(arr)-d:], arr[:d])
	copy(arr, tempArr)
}

func RotateRight(arr []int, d int) {
	d %= len(arr)

	tempArr := make([]int, len(arr))
	copy(tempArr, arr[len(arr)-d:])
	copy(tempArr[d:], arr[:len(arr)-d])
	copy(arr, tempArr)
}
