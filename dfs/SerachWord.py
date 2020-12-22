from typing import List


# 79. 单词搜索
# 给定一个二维网格和一个单词，找出该单词是否存在于网格中。
#
# 单词必须按照字母顺序，通过相邻的单元格内的字母构成，
# 其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        def backTrack(index, x, y, used):
            if board[x][y] != word[index]:
                return False
            else:
                if index == len(word) - 1:
                    return True
                used[x][y] = 1
                isFound = False
                for newX, newY in [(x, y - 1), (x, y + 1), (x - 1, y), (x + 1, y)]:
                    if 0 <= newX < len(board) and 0 <= newY < len(board[0]) and used[newX][newY] == 0:
                        isFound = backTrack(index + 1, newX, newY, used)
                    if isFound:
                        return True
                used[x][y] = 0
                return False

        isFound = False
        m = len(board)
        n = len(board[0])
        used = [[0] * n for _ in range(m)]
        for i in range(m):
            for j in range(n):
                if backTrack(0, i, j, used):
                    return True
        return False
