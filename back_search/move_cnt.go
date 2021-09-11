package backtrack_search

func movingCount(m int, n int, k int) int {
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	return dfs(0, 0, m, n, k, visited)
}

func dfs(i, j, m, n, k int, visited [][]bool) int {
	if i < 0 || j < 0 || i >= m || j >= n || (i/10+i%10+j/10+j%10) > k || visited[i][j] {
		return 0
	}
	visited[i][j] = true
	return 1 + dfs(i+1, j, m, n, k, visited) + dfs(i, j+1, m, n, k, visited) + dfs(i-1, j, m, n, k, visited) + dfs(i, j-1, m, n, k, visited)
}
