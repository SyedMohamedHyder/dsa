package arrays

// Time Complexity: O(log n)
// Space Complexity: O(1)
func BinarySearchLoop(arr []int, target int) int {

	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else if arr[mid] > target {
			high = mid - 1
		}
	}

	return -1
}

// Time Complexity: O(log n)
// Space Complexity: O(n)
func BinarySearchRecurse(arr []int, target, low, high int) int {

	if low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			return BinarySearchRecurse(arr, target, mid+1, high)
		} else if arr[mid] > target {
			return BinarySearchRecurse(arr, target, low, mid-1)
		}
	}

	return -1
}
