# 39. 组合总和
# 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
#
# candidates 中的数字可以无限制重复被选取。
#
# 说明：
#
# 所有数字（包括 target）都是正整数。
# 解集不能包含重复的组合。

import copy
from typing import List


import copy
class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        def backTrack(ans,ansList,candidates,target):
            if target<0:
                return
            if target==0:
                ansCopy=copy.deepcopy(ans)
                ansCopy.sort()
                if ansCopy not in ansList:
                    ansList.append(ansCopy)
                return
            # 递归搜索
            for number in candidates:
                ans.append(number)
                backTrack(ans,ansList,candidates,target-number)
                ans.pop()
        ans=[]
        ansList=[]
        backTrack(ans,ansList,candidates,target)
        return ansList
