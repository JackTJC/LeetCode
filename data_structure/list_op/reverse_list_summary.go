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
	ret := reverseListRecur(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}

//反转链表2

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var start *ListNode
	fakeHead := &ListNode{
		Next: head,
	}
	var preOfStart *ListNode
	cnt := 0
	start = fakeHead
	for cnt != left {
		preOfStart = start
		start = start.Next
		cnt++
	}
	reverseHead := start
	var pre *ListNode
	for cnt != right+1 {
		temp := start.Next
		start.Next = pre
		pre = start
		start = temp
		cnt++
	}
	preOfStart.Next = pre
	reverseHead.Next = start
	return fakeHead.Next
}
