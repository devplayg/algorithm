package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	size := 1000
	sorted := make([]int, size)

	for i := 0; i < size; i++ {
		sorted[i] = 2 * i
	}

	for i := 0; i < size; i++ {
		index := binarySearch(sorted, 2*i)
		if index != i {
			t.Error()
		}
	}

	if binarySearch(sorted, 3) != -1 {
		t.Error()
	}
}
