package tree

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	var cutIdx int
	for idx := range inorder {
		if inorder[idx] == preorder[0] {
			cutIdx = idx
		}
	}
	li, ri := inorder[:cutIdx], inorder[cutIdx+1:]
	lp, rp := preorder[1:1+len(li)], preorder[len(li)+1:]
	ret := &TreeNode{
		Val: preorder[0],
	}
	ret.Left = buildTree(lp, li)
	ret.Right = buildTree(rp, ri)
	return ret
}
