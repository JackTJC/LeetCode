package backtrack_search

func findNumberIn2DArray(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	//先在第一行找
	left, right := 0, n-1
	for left < right {
		mid := (left + right) / 2
		if matrix[0][mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	searchMax := right
	for i := 0; i <= searchMax; i++ {
		top, bottom := 0, m-1
		for top <= bottom {
			mid := (top + bottom) / 2
			if matrix[mid][i] < target {
				top = mid + 1
			} else if matrix[mid][i] > target {
				bottom = mid - 1
			} else {
				return true
			}
		}
	}
	return false
}
