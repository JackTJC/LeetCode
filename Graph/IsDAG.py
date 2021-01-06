# 207. 课程表
# 你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。
#
# 在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
#
# 给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？
from typing import List


class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        # 入度hash表,保存了入度的值
        indegree = {i: 0 for i in range(numCourses)}
        # 邻接hash表，保存了所有与之邻接的节点
        neighbor = {i: list() for i in range(numCourses)}
        for x, y in prerequisites:
            indegree[x] += 1
            neighbor[y].append(x)
        while len(indegree) > 0:
            temp = []
            for k, v in indegree.items():
                if v == 0:  # 找到度为0的节点，对其邻接节点的度减1
                    temp.append(k)
                    for n in neighbor[k]:
                        indegree[n] -= 1
            # 一轮遍历发现已经没有度为0的节点，返回false
            if len(temp) == 0:
                return False
            # 删除所有度为0的节点
            for t in temp:
                del indegree[t]
        return True
