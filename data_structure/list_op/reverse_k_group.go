package list_op

/*
* 25. K 个一组翻转链表
* 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
* k 是一个正整数，它的值小于或等于链表的长度。
* 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
*
* 进阶：
* 你可以设计一个只使用常数额外空间的算法来解决此问题吗？
* 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	cnt := 0
	tmp := head
	for tmp != nil {
		cnt++
		tmp = tmp.Next
	}
	var l, r, nxt, ret *ListNode
	nxt = head
	for i := 0; i < cnt/k; i++ {
		preR := r
		l, r, nxt = reverse(nxt, k, 1)
		if i == 0 {
			ret = l
			continue
		}
		preR.Next = l
	}
	if nxt != nil {
		r.Next = nxt
	}
	return ret
}

// 魔改后的reverse函数
// @head: 链表头
// @k: 头部需要翻转的k个节点
// @cur: 输入为0
//
// @return:
// 前k个反转的头节点、尾节点，后续没有反转链表的头节点
//
// example:
// 输入
// 1->2->3->4->5 3 0
// 输出
// 3 -> 2-> 1 		 4->5
// |        | 		 |
// ret1   ret2      ret3
func reverse(head *ListNode, k, cur int) (*ListNode, *ListNode, *ListNode) {
	if head == nil || head.Next == nil || cur >= k {
		return head, head, head.Next
	}
	l, _, nxt := reverse(head.Next, k, cur+1)
	head.Next.Next = head
	head.Next = nil
	return l, head, nxt
}
