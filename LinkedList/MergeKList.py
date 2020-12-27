# 23. 合并K个升序链表
# 给你一个链表数组，每个链表都已经按升序排列。
#
# 请你将所有链表合并到一个升序链表中，返回合并后的链表。
from typing import List
# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:
    def mergeKLists(self, lists: List[ListNode]) -> ListNode:
        def mergeTwoList(node1,node2):
            dummyHead=ListNode(-2**31)
            cur=dummyHead
            p1=node1
            p2=node2
            while p1 and p2:
                if p1.val <= p2.val:
                    cur.next=p1
                    p1=p1.next
                else:
                    cur.next=p2
                    p2=p2.next
                cur=cur.next
            if p1:
                cur.next=p1
            if p2:
                cur.next=p2
            return dummyHead.next
        while True:
            temp=[]
            while len(lists)>1:
                temp.append(mergeTwoList(lists.pop(),lists.pop()))
            if lists:
                temp.append(lists.pop())
            if len(temp)==0:
                return None
            if len(temp)==1:
                return temp[0]
            lists=temp[:]
