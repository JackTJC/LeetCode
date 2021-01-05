from typing import List


# 48. 旋转图像
# 给定一个 n × n 的二维矩阵表示一个图像。
#
# 将图像顺时针旋转 90 度。
#
# 说明：
#
# 你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。

class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """

        # 旋转一个格子
        def helper(mat, startpos, num):
            temp = mat[startpos][startpos]
            for i in range(num):
                mat[startpos + i][startpos] = mat[startpos + i + 1][startpos]
            for i in range(num):
                mat[startpos + num][i] = mat[startpos + num][i + 1]
            for i in range(num):
                mat[startpos + num - i][num] = mat[startpos + num - i - 1][num]
            for i in range(num):
                if i == num - 1:
                    mat[startpos][num - i] = temp
                    break
                mat[startpos][num - i] = mat[startpos][num - i - 1]

        for i in range(int(len(matrix) / 2)):
            n = len(matrix)
            for j in range(n - 2 * i - 1):
                helper(matrix, i, n - 2 * i - 1)


if __name__ == '__main__':
    Solution.rotate(None, [[5, 1, 9, 11], [2, 4, 8, 10], [13, 3, 6, 7], [15, 14, 12, 16]])
