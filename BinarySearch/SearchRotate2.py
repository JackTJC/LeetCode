# 81. 搜索旋转排序数组 II
# 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
#
# ( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。
#
# 编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。
from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> bool:
        while len(nums) > 1 and nums[0] == nums[-1]:
            nums.pop()
        n = len(nums)
        begin = 0
        end = n - 1
        while begin <= end:
            mid = (begin + end) // 2
            if nums[mid] == target:
                return True
            if nums[0] <= nums[mid]:
                if nums[0] <= target < nums[mid]:
                    end = mid - 1
                else:
                    begin = mid + 1
            elif nums[0] > nums[mid]:
                if nums[mid] < target <= nums[-1]:
                    begin = mid + 1
                else:
                    end = mid - 1
        return False
