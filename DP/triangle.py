from typing import List
#120

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
