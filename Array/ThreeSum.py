class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        first = 0
        ans = []
        for first in range(len(nums)):
            third = len(nums) - 1
            if first > 0 and nums[first] == nums[first - 1]:
                continue  # 跳过相同的
            target = -nums[first]
            for second in range(first + 1, len(nums)):
                if second > first + 1 and nums[second] == nums[second - 1]:
                    continue  # 跳过相同的
                while second < third and nums[second] + nums[third] > target:
                    third -= 1
                if third == second:
                    break
                if nums[second] + nums[third] == target:
                    ans.append([nums[first], nums[second], nums[third]])
        return ans

