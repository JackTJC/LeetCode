# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        ans = ListNode()
        pointer = ans
        Ci = 0
        while l1 is not None or l2 is not None:
            twoSum = (l1.val if l1 else 0) + (l2.val if l2 else 0) + Ci
            pointer.next = ListNode(val=twoSum % 10)
            Ci = 1 if twoSum >= 10 else 0
            pointer = pointer.next
            l1 = l1.next if l1 else None
            l2 = l2.next if l2 else None
        if Ci == 1:
            pointer.next = ListNode(val=1)
        return ans.next
