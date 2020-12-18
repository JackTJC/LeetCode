import copy
from typing import List

class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        nums.sort()

        def backTrack(ans, ansList, index):
            if len(ans) == 4:
                if sum(ans) == target:
                    ansList.append(copy.deepcopy(ans))
                return
            for i in range(index, len(nums)):
                ans.append(nums[index])
                backTrack(ans, ansList, index + 1)
                ans.pop()

        ans = []
        ansList = []
        backTrack(ans, ansList, 0)
        return ansList

if __name__ == '__main__':
    Solution.fourSum(None,[1,0,-1,0,-2,2],0)