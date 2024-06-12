package arrays

func LeftShiftArr(arr []int, d int) {
	for i := 0; i < d; i++ {
		for j := 0; j < len(arr)-1; j++ {
			arr[j] = arr[j+1]
		}

		arr[len(arr)-1] = 0
	}

}

func RightShiftArr(arr []int, d int) {
	for i := 0; i < d; i++ {
		for j := len(arr) - 1; j > 0; j-- {
			arr[j] = arr[j-1]
		}

		arr[0] = 0
	}
}
