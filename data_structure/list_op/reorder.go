package list_op

func reorderList(head *ListNode) {
	fast, slow := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var midHead *ListNode
	if fast == nil {
		midHead = rvList(slow)
	} else {
		midHead = rvList(slow.Next)
	}
	//merge
	for midHead != nil {
		t := head.Next
		head.Next = midHead
		midHead = midHead.Next
		head.Next.Next = t
		head = head.Next.Next
	}
	head.Next = nil
}

func rvList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	rvHead := rvList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return rvHead
}
