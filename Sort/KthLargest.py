# 215. 数组中的第K个最大元素
# 在未排序的数组中找到第 k 个最大的元素。请注意，
# 你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
from typing import List


class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        def adjustHeap(heap, n, i):
            """
            堆调整算法
            :param heap: 待调整的堆
            :param n: 堆的大小
            :param i: 开始调整的位置
            :return: None
            """
            k = 2 * i + 1
            temp = heap[i]
            while k < n:
                if k + 1 < n and heap[k + 1] < heap[k]:
                    k += 1
                if heap[k] < temp:
                    heap[i] = heap[k]
                    i = k
                else:
                    break
                k = 2 * k + 1
            heap[i] = temp

        # 维护一个小顶堆，输出堆顶元素
        n = len(nums)
        for i in range((k // 2) - 1, -1, -1):
            adjustHeap(nums, k, i)  # 建立堆时的调整
        for i in range(k, n):
            if nums[i] > nums[0]:
                nums[0] = nums[i]
                adjustHeap(nums, k, 0)  # 替换堆顶后的调整
        return nums[0]
