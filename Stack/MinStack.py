# 155. 最小栈
# 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
#
# push(x) —— 将元素 x 推入栈中。
# pop() —— 删除栈顶的元素。
# top() —— 获取栈顶元素。
# getMin() —— 检索栈中的最小元素。


#辅助栈算法
class MinStack:

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.stack=list()
        self.helpStack=[2**31-1,]


    def push(self, x: int) -> None:
        self.stack.append(x)
        if x<self.helpStack[-1]:
            self.helpStack.append(x)
        else:
            self.helpStack.append(self.helpStack[-1])

    def pop(self) -> None:
        self.stack.pop()
        self.helpStack.pop()


    def top(self) -> int:
        return self.stack[-1]


    def getMin(self) -> int:
        return self.helpStack[-1]