from typing import List


class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        def backTrack(index, x, y, used):
            if board[x][y] != word[index]:
                return False
            else:
                if index==len(word)-1:
                    return True
                used[x][y] = 1
                isFound = False
                for newX, newY in [(x, y - 1), (x, y + 1), (x - 1, y), (x + 1, y)]:
                    if 0 <= newX < len(board) and 0 <= newY < len(board[0]) and used[newX][newY] == 0:
                        isFound = backTrack(index + 1, newX, newY, used)
                    if isFound:
                        return True
                return False

        isFound = False
        m = len(board)
        n = len(board[0])
        used = [[0] * n for _ in range(m)]
        for i in range(m):
            for j in range(n):
                isFound = backTrack(0, i, j, used)
                if isFound:
                    return True
        return False


if __name__ == '__main__':
    Solution.exist(None,[["C","A","A"],["A","A","A"],["B","C","D"]],"AAB")