from typing import List
class Solution:
    def smallestK(self, arr: List[int], k: int) -> List[int]:
        # topK问题，使用堆排序解决
        def adjustHead(heap, i, length):
            temp = heap[i]
            j = 2 * i + 1
            while j < length:
                if j + 1 < length and heap[j] > heap[j + 1]:
                    j += 1
                if heap[j] < temp:
                    heap[i] = heap[j]
                    i = j
                else:
                    break
                j = 2 * j + 1
            heap[i] = temp

        # 维护一个大小为k的小顶堆
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

if __name__ == '__main__':
    Solution.smallestK(None,[9,8,7,6,5,4,3,2,1],4)