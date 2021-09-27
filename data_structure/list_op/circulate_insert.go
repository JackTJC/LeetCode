package list_op

/*
 * 剑指 Offer II 029. 排序的循环链表
 * 给定循环升序列表中的一个点，写一个函数向这个列表中插入一个新元素 insertVal ，使这个列表仍然是循环升序的。
 * 给定的可以是这个列表中任意一个顶点的指针，并不一定是这个列表中最小元素的指针。
 * 如果有多个满足条件的插入位置，可以选择任意一个位置插入新的值，插入后整个列表仍然保持有序。
 * 如果列表为空（给定的节点是 null），需要创建一个循环有序列表并返回这个节点。否则。请返回原先给定的节点。
 */
func insert(aNode *Node, x int) *Node {
	pp := aNode
	nNode := &Node{
		Val: x,
	}
	if pp == nil {
		nNode.Next = nNode
		return nNode
	}
	p := aNode.Next
	if p == pp {
		aNode.Next = nNode
		nNode.Next = aNode
		return aNode
	}
	canBreak := func(x1, x2, x int) bool {
		//在中间某个位置
		if x1 <= x && x <= x2 {
			return true
		}
		//在开头
		if x >= x1 && x >= x2 && x1 > x2 {
			return true
		}
		//在结尾
		if x <= x1 && x <= x2 && x2 < x1 {
			return true
		}
		return false
	}
	for {
		if canBreak(pp.Val, p.Val, x) || p == aNode { //p==aNode这种情况就是所有节点值相同
			break
		}
		pp = p
		p = p.Next
	}
	pp.Next = nNode
	nNode.Next = p
	return aNode
}
