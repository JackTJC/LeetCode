package list_op

/**
 * Definition for a Node.
 */
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

type copier struct {
	copied map[*Node]*Node
}

func (c *copier) run(head *Node) *Node {
	if head == nil {
		return nil
	}
	var ret *Node
	if _, ok := c.copied[head]; !ok {
		ret = &Node{
			Val: head.Val,
		}
		c.copied[head] = ret
		ret.Next = c.run(head.Next)
		ret.Random = c.run(head.Random)
	}
	return c.copied[head]
}

func copyRandomList(head *Node) *Node {
	c := &copier{
		copied: make(map[*Node]*Node),
	}
	return c.run(head)
}
