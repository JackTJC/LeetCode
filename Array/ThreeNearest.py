from typing import List
# 16. 最接近的三数之和
# 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。
# 找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
# 二分不适合这一类题目
class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        nums.sort()
        ans = 10 ** 7
        n = len(nums)
        for i in range(n):
            if i > 0 and nums[i] == nums[i - 1]:
                continue
            j, k = i + 1, n - 1
            while j < k:
                s = nums[i] + nums[j] + nums[k]
                if abs(ans - target) > abs(s - target):
                    ans = s
                if s == target:
                    return target
                if s > target:
                    k0 = k - 1
                    while j < k0 and nums[k0] == nums[k]:
                        k0 -= 1
                    k = k0
                else:
                    j0 = j + 1
                    while j0 < k and nums[j0] == nums[j]:
                        j0 += 1
                    j = j0
        return ans
