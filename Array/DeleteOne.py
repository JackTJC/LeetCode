from typing import List
# 27. 移除元素
# 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
#
# 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
#
# 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        if len(nums)==1 and nums[0]!=val:
            return 1
        begin=0
        end=len(nums)-1
        while begin<end:
            while begin<len(nums) and nums[begin]!=val:
                begin+=1
            while end>=0 and nums[end]==val:
                end-=1
            if begin<end:
                temp=nums[begin]
                nums[begin]=nums[end]
                nums[end]=temp
        return begin

if __name__ == '__main__':
    Solution.removeElement(None,[1,1,1,2,2,2,3,3,3,1,1,1],1)