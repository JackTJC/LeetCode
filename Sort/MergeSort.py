from typing import List


# 4. 寻找两个正序数组的中位数
# 给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。
#
# 进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？

class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        array = [0] * (len(nums1) + len(nums2))
        if len(array) == 0:
            return 0
        pointer1 = 0
        pointer2 = 0
        i = 0
        # 归并排序
        while pointer1 <= len(nums1) - 1 and pointer2 <= len(nums2) - 1:
            if nums1[pointer1] <= nums2[pointer2]:
                array[i] = nums1[pointer1]
                pointer1 += 1
            else:
                array[i] = nums2[pointer2]
                pointer2 += 1
            i += 1
        while pointer1 <= len(nums1) - 1:
            array[i] = nums1[pointer1]
            i += 1
            pointer1 += 1
        while pointer2 <= len(nums2) - 1:
            array[i] = nums2[pointer2]
            i += 1
            pointer2 += 1
        if len(array) % 2 == 0:
            return (array[int(len(array) / 2) - 1] + array[int(len(array) / 2)]) / 2
        else:
            return array[int(len(array) / 2)]
