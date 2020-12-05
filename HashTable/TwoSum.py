from typing import List
# 1. 两数之和
# 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
#
# 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hashMap={}
        for i in range(len(nums)):
            if hashMap.get(nums[i]):
                hashMap[nums[i]].append(i)
            else:
                hashMap[nums[i]]=[i]
        for num in nums:
            val=hashMap.get(target-num)
            if val is not None:
                if num==target-num:
                    if len(val)>1:
                        return val[:2]
                    else:
                        continue
                else:
                    return [hashMap[num][0],val[0]]