import copy
from typing import List
class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        def backTrack(target,ans,ansList,selList):
            if target<0:
                return
            if target==0:
                ansCopy=copy.deepcopy(ans)
                ansCopy.sort()
                if ansCopy not in ansList:
                    ansList.append(ansCopy)
                return
            for i in range(len(candidates)):
                if selList[i]==0:
                    #select
                    ans.append(candidates[i])
                    selList[i]=1
                    backTrack(target-candidates[i],ans,ansList,selList)
                    ans.pop()
                    selList[i]=0
        if sum(candidates)<target:
            return []
        selectedList=[0]*len(candidates)
        ans=[]
        ansList=[]
        backTrack(target,ans,ansList,selectedList)
        return ansList

if __name__ == '__main__':
    Solution.combinationSum2(None,[1],1)