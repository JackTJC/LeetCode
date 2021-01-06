# 189. 旋转数组
# 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
from typing import List


class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """

        def reverse(nums, start, end):
            while start < end:
                temp = nums[start]
                nums[start] = nums[end]
                nums[end] = temp
                start += 1
                end -= 1
        # 整体翻转
        nums.reverse()
        k = k % len(nums)
        # 翻转前K个
        reverse(nums, 0, k - 1)
        # 翻转后面的部分
        reverse(nums, k, len(nums) - 1)
