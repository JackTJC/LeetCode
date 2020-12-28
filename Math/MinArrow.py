from typing import List


class Solution:
    def findMinArrowShots(self, points: List[List[int]]) -> int:
        sorter = lambda x: x[0] * 10 + x[1]
        points.sort(key=sorter)
        needArrow = []
        needArrow.append(points[0])
        for i in range(1, len(points)):
            if points[i][0] <= needArrow[-1][1]:
                end = needArrow.pop()
                needArrow.append([points[i][0], end[1]])
            else:
                needArrow.append(points[i][0])
        return len(needArrow)


if __name__ == '__main__':
    Solution.findMinArrowShots(None, [[9, 12], [1, 10], [4, 11], [8, 12], [3, 9], [6, 9], [6, 7]])
