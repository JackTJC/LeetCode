package sort_alg

import (
	"testing"
)

func TestHeap(t *testing.T) {
	arr := []int{
		10, 9, 8, 7, 6, 5, 4, 3, 2, 1,
	}
	HeapSort(arr)
	t.Log(arr)
}
