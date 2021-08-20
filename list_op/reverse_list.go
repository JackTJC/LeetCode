package list_op

func reverseList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	ret := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}
