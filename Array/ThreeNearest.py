from typing import List
# 二分不适合这一类题目
class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        nums.sort()
        ans=10**7
        n=len(nums)
        for i in range(n):
            if i>0 and nums[i]==nums[i-1]:
                continue
            j,k=i+1,n-1
            while j<k:
                s=nums[i]+nums[j]+nums[k]
                if abs(ans-target)>abs(s-target):
                    ans=s
                if s == target:
                    return target
                if s>target:
                    k0=k-1
                    while j<k0 and nums[k0]==nums[k]:
                        k0-=1
                    k=k0
                else:
                    j0=j+1
                    while j0<k and nums[j0]==nums[j]:
                        j0+=1
                    j=j0
            return ans