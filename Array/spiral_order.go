package array

/*
*54. 螺旋矩阵
*给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
 */

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	return spiralHelper(matrix, 0, 0, m, n)
}

// row,col 新的起始位置
// m,n 新的矩阵大小
func spiralHelper(matrix [][]int, row, col, m, n int) []int {
	ret := make([]int, 0)
	if m <= 0 || n <= 0 {
		return ret
	}
	/*
	*	把边界分为4部分，top，right，bottom，left
	*	分别求和
	 */

	// top
	for i := 0; i < n; i++ {
		ret = append(ret, matrix[row][col+i])
	}
	if m == 1 {
		return ret
	}
	//right
	for i := 1; i < m; i++ {
		ret = append(ret, matrix[row+i][col+n-1])
	}
	if n == 1 {
		return ret
	}
	//bottom
	for i := n - 2; i >= 0; i-- {
		ret = append(ret, matrix[row+m-1][col+i])
	}
	//left
	for i := m - 2; i >= 1; i-- {
		ret = append(ret, matrix[row+i][col])
	}
	ret = append(ret, spiralHelper(matrix, row+1, col+1, m-2, n-2)...)
	return ret
}
