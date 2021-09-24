package list_op

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}
	slow := head
	var pSlow *ListNode
	for fast != nil {
		fast = fast.Next
		pSlow = slow
		slow = slow.Next
	}
	pSlow.Next = slow.Next
	return head
}
