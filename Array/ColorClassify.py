from typing import List
class Solution:
    def sortColors(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        i = 0
        while i < len(nums) - 1:
            if nums[i] == 0:
                # 插入首部
                for j in range(i, 0, -1):
                    nums[j] = nums[j - 1]
                nums[0] = 0
            elif nums[i] == 2:
                # 插入尾部
                if min(nums[i:]) == 2:
                    return
                for j in range(i, len(nums) - 1):
                    nums[j] = nums[j + 1]
                nums[-1] = 2
                i -= 1
            i += 1
if __name__ == '__main__':
    Solution.sortColors(None,[1,0])