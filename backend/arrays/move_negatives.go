package arrays

func ArrangeNegatives(arr []int) {
	l, m := 0, len(arr)-1
	for l < m {
		for arr[l] < 0 {
			l++
		}

		for arr[m] >= 0 {
			m--
		}

		if l < m {
			arr[l], arr[m] = arr[m], arr[l]
		}
	}
}
