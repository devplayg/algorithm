package search

func binarySearch(arr []int, item int) int {
	start, end := 0, len(arr)-1

	for start <= end {
		mid := ((end - start) >> 1) + start

		if arr[mid] == item {
			return mid
		} else if arr[mid] < item {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}
