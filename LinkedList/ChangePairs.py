# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


# 24.
# 两两交换链表中的节点
# 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
#
# 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

class Solution:
    def swapPairs(self, head: ListNode) -> ListNode:
        newHead = ListNode(0, head)
        temp = newHead
        # 使用新的头节点方便终止条件判断
        while temp.next and temp.next.next:
            node1 = temp.next
            node2 = temp.next.next
            temp.next = node2
            node1.next = node2.next #连接上
            node2.next = node1
            temp = node1
        return newHead.next


if __name__ == '__main__':
    a=ListNode(1)
    a.next=ListNode(2)
    a.next.next=ListNode(3)
    a.next.next.next=ListNode(4)
    Solution.swapPairs(None,a)