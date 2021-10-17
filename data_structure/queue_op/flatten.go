package queueop

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

/*
*  剑指 Offer II 028. 展平多级双向链表
*  多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。
*  这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。
*  给定位于列表第一级的头节点，请扁平化列表，即将这样的多级双向链表展平成普通的双向链表，使所有结点出现在单级双链表中
 */
func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	stack := make([]*Node, 0)
	stack = append(stack, root)
	// 用一个duplicate的头方便操作
	dupHead := &Node{}
	pHead := dupHead
	for len(stack) > 0 {
		//模拟栈的pop操作
		popV := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		dupHead.Next = popV
		popV.Prev = dupHead
		dupHead = dupHead.Next
		if popV.Next != nil {
			stack = append(stack, popV.Next)
			popV.Next = nil
		}
		if popV.Child != nil {
			stack = append(stack, popV.Child)
			popV.Child = nil
		}
	}
	//头节点的pre指针为空
	pHead.Next.Prev = nil
	return pHead.Next
}
