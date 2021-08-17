# 155. 最小栈
# 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
#
# push(x) —— 将元素 x 推入栈中。
# pop() —— 删除栈顶的元素。
# top() —— 获取栈顶元素。
# getMin() —— 检索栈中的最小元素。
# 保留差值算法

class MinStack:

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.diffStack=list()
        self.minVal=-1


    def push(self, x: int) -> None:
        if not self.diffStack:
            self.diffStack.append(0)
            self.minVal=x
        else:
            diff=x-self.minVal
            self.diffStack.append(diff)
            self.minVal=self.minVal if self.minVal<x else x

    def pop(self) -> None:
        top=self.diffStack.pop()
        self.minVal=self.minVal if top>0 else self.minVal-top


    def top(self) -> int:
        return self.minVal if self.diffStack[-1]<0 else self.diffStack[-1]+self.minVal


    def getMin(self) -> int:
        return self.minVal