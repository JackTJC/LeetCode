# 0804
from typing import List


class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        def backTrack(nodeList, i, arr, ans):
            ans.append([j for j in nodeList])
            for i in range(i, len(arr)):
                nodeList.append(arr[i])
                backTrack(nodeList, i + 1, arr, ans)
                nodeList.pop()

        ans = []
        nodeList = []
        backTrack(nodeList, 0, nums, ans)
        return ans
