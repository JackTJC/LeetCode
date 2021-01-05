# 154. 寻找旋转排序数组中的最小值 II
# 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
#
# ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
#
# 请找出其中最小的元素。
#
# 注意数组中可能存在重复的元素。
from typing import List
class Solution:
    def findMin(self, nums: List[int]) -> int:
        n=len(nums)
        begin=0
        end=n-1
        while begin<end:
            mid=(begin+end)//2
            if nums[mid]<nums[end]:
                end=mid
            elif nums[mid]>nums[end]:
                begin=mid+1
            else:
                end-=1
        return nums[begin]