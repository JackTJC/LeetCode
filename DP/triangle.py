# 120. 三角形最小路径和
# 给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
#
# 相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
from typing import List

class Solution:
    def minimumTotal(self, triangle: List[List[int]]) -> int:
        res = list()
        for i in range(len(triangle)):
            row = list()
            if i == 0:
                row = triangle[i]
            else:
                for j in range(len(triangle[i])):
                    if j == 0:
                        row.append(triangle[i][0] + res[i - 1][0])
                    elif j == len(triangle[i]) - 1:
                        row.append(triangle[i][j] + res[i - 1][j - 1])
                    else:
                        row.append(triangle[i][j] + min(res[i - 1][j], res[i - 1][j - 1]))
            res.append(row)
        return min(res[len(res) - 1])
