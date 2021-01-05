# 213. 打家劫舍 II
# 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，
# 这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，
# 如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
#
# 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，能够偷窃到的最高金额。
from typing import List


class Solution:
    def rob(self, nums: List[int]) -> int:
        n = len(nums)
        if n <= 3:
            return max(nums)
        #  偷窃最后一个房子
        today = yesterday = 0
        for i in range(1, n):
            yesterday, today = today, max(yesterday + nums[i], today)
        t1 = today
        # 不偷窃最后一个房子
        today = yesterday = 0
        for i in range(n - 1):
            yesterday, today = today, max(yesterday + nums[i], today)
        t2 = today
        return max(t1, t2)
