# 154. 寻找旋转排序数组中的最小值 II
# 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
#
# ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
#
# 请找出其中最小的元素。
#
# 注意数组中可能存在重复的元素。
from typing import List


class Solution:
    def findMin(self, nums: List[int]) -> int:
        n = len(nums)
        begin = 0
        end = n - 1
        while begin < end:
            mid = (begin + end) // 2
            if nums[mid] < nums[end]:  # mid比end小,必定在mid及之前
                end = mid
            elif nums[mid] > nums[end]:  # mid比end大,必定在mid加一及以后
                begin = mid + 1
            else:  # 不敢确定只能将end-1
                end -= 1
        return nums[begin]

# 解法2 GO语言
#
# func findMin(nums []int) int {
# 	n:=len(nums)
# 	l:=0
# 	r:=n-1
# 	for l<r{
# 		m:=int((l+r)/2)
#       //对于任何一个旋转数组，除非没有旋转，否则最左边必然大于等于最右边
# 		if nums[l]<nums[r]{
# 			return nums[l]
# 		}
# 		if nums[m]>nums[l] {
#           //左边有序，且排除顺序的情况，必然向右边搜索
# 			l=m+1
# 		}else if nums[m]<nums[l] {
#           //右边有序,必然向左边搜索
# 			r=m
# 		}else {
#           //如果nums[mid] == nums[l]并不能区分是左边都一样还是右边都一样，需要
# 			l++
# 		}
# 	}
# 	return nums[l]
# }
