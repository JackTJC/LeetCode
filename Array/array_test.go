package array

import "testing"

func TestFR(t *testing.T) {
	t.Log(findRepeatNumber([]int{3, 2, 1, 0, 0}))
}

func TestSearch(t *testing.T) {
	arr := []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 7, 8, 9, 10}
	t.Log(search(arr, 1))
	t.Log(search(arr, 2))
	t.Log(search(arr, 3))
	t.Log(search(arr, 4))
	t.Log(search(arr, 5))
	t.Log(search(arr, 6))
	t.Log(search(arr, 7))
	t.Log(search(arr, 8))
	t.Log(search(arr, 9))
	t.Log(search(arr, 11))
	t.Log(search(arr, 0))
}

func TestMN(t *testing.T) {
	t.Log(missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 7, 9}))
	t.Log(missingNumber([]int{0, 2}))
	t.Log(missingNumber([]int{0, 1, 3}))
	t.Log(missingNumber([]int{0}))
	t.Log(missingNumber([]int{1}))
	t.Log(missingNumber([]int{1, 2}))
}
