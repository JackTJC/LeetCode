# 6. Z 字形变换
# 将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
#
# 比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
#
# L   C   I   R
# E T O E S I I G
# E   D   H   N
# 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
#
# 请你实现这个将字符串进行指定行数变换的函数：
#
# string convert(string s, int numRows);
class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows==1:
            return s
        display=[]
        downForward=True
        for _ in range(numRows):
            display.append([0])
        pos=[0,0]
        count=0
        for c in s:
            if downForward:
                display[pos[0]][pos[1]]=c
                count+=1
                pos[0]+=1
            else:
                display[pos[0]][pos[1]]=c
                count+=1
                pos[0]-=1
                pos[1]+=1
                for col in display:#扩充1列
                    col.append(0)
            if count==numRows-1:
                downForward=not downForward
                for col in display:#扩充1列
                    col.append(0)
                count=0
        ans=""
        for row in display:
            for c in row:
                if c!=0:
                    ans+=c
        return ans