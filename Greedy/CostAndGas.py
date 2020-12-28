# https://leetcode-cn.com/problems/gas-station/solution/shi-yong-tu-de-si-xiang-fen-xi-gai-wen-ti-by-cyayc/
# 134. 加油站
# 在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
#
# 你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
#
# 如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。
from typing import List


class Solution:
    def canCompleteCircuit(self, gas: List[int], cost: List[int]) -> int:
        n = len(gas)
        total = 0
        minVal = 2 ** 31 - 1
        minIndex = 0
        for i in range(n):
            total += gas[i] - cost[i]
            if total < minVal:
                minVal = total
                minIndex = i
        return -1 if total < 0 else (minIndex + 1) % n
