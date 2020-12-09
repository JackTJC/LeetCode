from typing import List

# 22.
# 括号生成
# 数字
# n
# 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且
# 有效的
# 括号组合。

class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        def backTrack(leftNum,rightNum,ans,ansList,length):
            if rightNum>leftNum or leftNum>length:
                return
            if len(ans)==length*2 and rightNum==leftNum:
                ansList.append(''.join(ans))
                return
            ans.append('(')
            backTrack(leftNum+1,rightNum,ans,ansList,length)
            ans.pop()
            ans.append(')')
            backTrack(leftNum,rightNum+1,ans,ansList,length)
            ans.pop()
            return
        if n==0:
            return [""]
        ans=['(']
        ansList=[]
        backTrack(1,0,ans,ansList,n)
        return ansList