package backtrack_search

import "testing"

func TestMinArray(t *testing.T) {
	t.Log(minArray([]int{2, 2, 2, 0, 1}))
	t.Log(minArray([]int{1, 0}))
	t.Log(minArray([]int{5, 6, 7, 8, 1, 2, 3, 4}))
	t.Log(minArray([]int{3, 3, 1, 3}))
}

func TestSearch2D(t *testing.T) {
	arr := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	t.Log(findNumberIn2DArray(arr, 5))
	t.Log(findNumberIn2DArray(arr, 30))
	t.Log(findNumberIn2DArray(arr, 23))
	t.Log(findNumberIn2DArray(arr, 14))
	t.Log(findNumberIn2DArray(arr, 27))
	t.Log(findNumberIn2DArray(arr, 15))
	t.Log(findNumberIn2DArray(arr, 18))
}

func TestComSum(t *testing.T) {
	combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
}
