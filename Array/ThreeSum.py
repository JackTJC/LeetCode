from typing import List


class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        if len(nums) < 3:
            return []

        def binarySearch(nums, begin, end, target):
            while begin <= end:
                mid = int((begin + end) / 2)
                if nums[mid] < target:
                    begin = mid + 1
                elif nums[mid] > target:
                    end = mid - 1
                else:
                    return True
            return False

        ans = []
        nums.sort()
        left = 0
        right = len(nums) - 1
        while left < right:
            target = -(nums[left] + nums[right])
            if nums[left + 1] <= target <= nums[right - 1]:
                res = binarySearch(nums, left + 1, right - 1, target)
                if res:
                    ans.append([nums[left], target, nums[right]])
                while nums[left] == nums[left + 1] and nums[right - 1] == nums[right] and left < right:
                    left += 1
                    right -= 1

            else:
                if target < nums[left + 1]:
                    right -= 1
                if target > nums[right - 1]:
                    left += 1
        return ans


if __name__ == '__main__':
    Solution.threeSum(None, [1, 2, -2, -1])
