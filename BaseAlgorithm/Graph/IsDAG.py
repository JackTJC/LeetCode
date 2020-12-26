from typing import List


class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        # 入度hash表
        indegree = {i: 0 for i in range(numCourses)}
        # 邻接hash表
        neighbor = {i: list() for i in range(numCourses)}
        for x, y in prerequisites:
            indegree[x] += 1
            neighbor[y].append(x)
        while len(indegree) > 0:
            temp = []
            for k, v in indegree.items():
                if v == 0:
                    temp.append(k)
                    for n in neighbor[k]:
                        indegree[n] -= 1
            if len(temp) == 0:
                return False
            for t in temp:
                del indegree[t]
        return True
