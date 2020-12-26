from typing import List
class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        def dfs(mark, i, j):
            nonlocal grid
            mark[i][j] = 1
            for (x, y) in [(i - 1, j), (i + 1, j), (i, j - 1), (i, j + 1)]:
                if 0 <= x < len(mark) and 0 <= y < len(mark[0]) and mark[x][y] == 0 and grid[x][y] == '1':
                    dfs(mark, x, y)

        m = len(grid)
        n = len(grid[0])
        mark = [[0] * n for _ in range(m)]
        total = 0
        for i in range(m):
            for j in range(n):
                if mark[i][j] == 0 and grid[i][j] == '1':
                    total += 1
                    dfs(mark, i, j)
        return total