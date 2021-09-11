from typing import List


class Solution:
    def leastInterval(self, tasks: List[str], n: int) -> int:
        if n == 0:
            return len(tasks)
        hashMap = [0] * 26
        for task in tasks:
            hashMap[ord(task) - ord('A')] += 1
        rowNeed = (max(hashMap) - 1) * (n + 1)
        return max((max(hashMap) - 1) * (n + 1) + hashMap.count(max(hashMap)), len(tasks))
