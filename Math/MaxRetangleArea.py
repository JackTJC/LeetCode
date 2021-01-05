# 223. 矩形面积
# 在二维平面上计算出两个由直线构成的矩形重叠后形成的总面积。
#
# 每个矩形由其左下顶点和右上顶点坐标表示，如图所示。

class Solution:
    def computeArea(self, A: int, B: int, C: int, D: int, E: int, F: int, G: int, H: int) -> int:
        def getArea(pos1, pos2):
            return abs(pos1[0] - pos2[0]) * abs(pos1[1] - pos2[1])
        # 核心为此处的求线段公共长度算法
        x = max(0, min(C, G) - max(A, E))
        y = max(0, min(D, H) - max(B, F))
        return getArea((A, B), (C, D)) + getArea((E, F), (G, H)) - x * y
