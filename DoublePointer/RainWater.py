# 42. 接雨水
# 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
from typing import List


class Solution:
    def trap(self, height: List[int]) -> int:
        # 双指针
        n=len(height)
        left=0
        right=n-1
        leftMax=0
        rightMax=0
        ans=0
        while left<right:
            leftMax=max(leftMax,height[left])
            rightMax=max(rightMax,height[right])
            if rightMax>=leftMax:
                ans+=leftMax-height[left]
                left+=1
            else:
                ans+=rightMax-height[right]
                right-=1
        return ans
