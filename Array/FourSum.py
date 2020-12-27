# 18. 四数之和
# 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，
# 使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
#
# 注意：
#
# 答案中不可以包含重复的四元组。


# 排序+双指针
from typing import List


class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        nums.sort()
        n = len(nums)
        ans = list()
        for first in range(0, n):
            if first > 0 and nums[first] == nums[first - 1]:
                continue
            for second in range(first + 1, n):
                if second > first + 1 and nums[second] == nums[second - 1]:
                    continue
                third = second + 1
                forth = n - 1
                while third < forth and third < n:
                    total = nums[first] + nums[second] + nums[third] + nums[forth]
                    if total == target:
                        ans.append([nums[first], nums[second], nums[third], nums[forth]])
                        while third < forth and nums[third + 1] == nums[third]:# 去重找到下一轮
                            third += 1
                        third += 1
                        while forth > third and nums[forth - 1] == nums[forth]:# 去重找到下一轮
                            forth -= 1
                        forth -= 1
                    elif total < target:
                        third += 1
                    else:
                        forth -= 1
        return ans
