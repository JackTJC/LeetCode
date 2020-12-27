# 25. K 个一组翻转链表
# 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
#
# k 是一个正整数，它的值小于或等于链表的长度。
#
# 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverseKGroup(self, head: ListNode, k: int) -> ListNode:
        def reverse(head):
            if not head.next:
                return head
            newHead = reverse(head.next)
            head.next.next = head
            head.next = None
            return newHead

        def helper(head):
            if not head:
                return None
            pHead = head
            pNext = None
            count = 1
            while head.next:
                if count == k:
                    pNext = head.next
                    head.next = None
                    break
                head = head.next
                count += 1
            if count < k:  # 不满足k个
                return pHead
            pHead = reverse(pHead)
            temp = pHead
            while temp.next:
                temp = temp.next
            temp.next = helper(pNext)
            return pHead

        return helper(head)
