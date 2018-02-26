package sort


func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[min] {
				arr[j], arr[min] = arr[min], arr[j]
				min = j
			}
		}
	}
}
