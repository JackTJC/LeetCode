package list_op

// 反转链表，迭代解法
func reverseListFor(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, next *ListNode
	pre, next = nil, head
	for next != nil {
		t := next.Next
		next.Next = pre
		pre = next
		next = t
	}
	return pre
}

//反转链表递归解法
func reverseListRecur(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ret := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}
