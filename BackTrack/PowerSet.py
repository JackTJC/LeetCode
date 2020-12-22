# 0804
# 面试题 08.04. 幂集
# 幂集。编写一种方法，返回某集合的所有子集。集合中不包含重复的元素。
#
# 说明：解集不能包含重复的子集。
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
