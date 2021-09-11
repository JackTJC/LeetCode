# 225. 用队列实现栈
# 使用队列实现栈的下列操作：

import collections


class MyStack:

    def __init__(self):
        self.q = collections.deque()  # append进队，popleft出队列

    def push(self, x: int) -> None:
        n = len(self.q)
        self.q.append(x)
        # 将原来在队列里的元素全部出队再入队即可
        for _ in range(n):
            self.q.append(self.q.popleft())

    def pop(self) -> int:
        return self.q.popleft()

    def top(self) -> int:
        return self.q[0]

    def empty(self) -> bool:
        return len(self.q) == 0
