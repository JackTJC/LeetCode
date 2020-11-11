# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None
class Solution:
    def partition(self, head: ListNode, x: int) -> ListNode:
        pLess=lessHead=ListNode(0)
        pOther=otherHead=ListNode(0)
        while head is not None:
            if head.val<x:
                pLess.next=ListNode(head.val)
                pLess=pLess.next
            else:
                pOther.next=ListNode(head.val)
                pOther=pOther.next
            head=head.next
        pLess.next=otherHead.next
        return lessHead.next