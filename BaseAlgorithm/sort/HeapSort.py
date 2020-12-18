from typing import List


class Solution:
    def smallestK(self, arr: List[int], k: int) -> List[int]:
        # topK问题，使用堆排序解决
        def adjustHead(heap, i, length):
            temp = heap[i]
            k = 2 * i + 1
            while k < length:
                if k + 1 < length and heap[k] < heap[k + 1]:
                    k += 1
                if heap[k] > temp:
                    heap[i] = heap[k]
                    i = k
                else:
                    break
                k = 2 * k + 1
            heap[i] = temp

        # 维护一个大小为k的大
        if k == 0:
            return []
        if k >= len(arr):
            return arr
        heap = arr[:k]
        for i in range(int(k / 2) - 1, -1, -1):
            adjustHead(heap, i, k)
        for i in range(k, len(arr)):
            if arr[i] < heap[0]:
                heap[0] = arr[i]
                adjustHead(heap, 0, k)
        return heap
