package arrays

func ReverseCopyBack(arr []int) {
	copyArr := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		copyArr[i] = arr[len(arr)-1-i]
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = copyArr[i]
	}
}

func ReverseSwap(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}
