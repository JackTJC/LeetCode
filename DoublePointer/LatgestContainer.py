# 11. 盛最多水的容器
# 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
# 在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。
# 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
from typing import List


class Solution:
    def maxArea(self, height: List[int]) -> int:
        right = len(height) - 1
        left = 0
        ans = 0
        while left < right:
            ans = max(ans, (right - left) * min(height[left], height[right]))
            if height[left] < height[right]:
                left += 1
            else:
                right -= 1
        return ans


if __name__ == '__main__':
    Solution.maxArea(None, [2, 3, 10, 5, 7, 8, 9])
