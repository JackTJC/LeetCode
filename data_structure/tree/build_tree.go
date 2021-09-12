package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。
假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	var splitIdx int
	for idx, num := range inorder {
		if num == rootVal {
			splitIdx = idx
		}
	}
	lIn, rIn := inorder[:splitIdx], inorder[splitIdx+1:]
	lPre, rPre := preorder[1:1+len(lIn)], preorder[1+len(lIn):]
	r := &TreeNode{
		Val:   rootVal,
		Left:  buildTree(lPre, lIn),
		Right: buildTree(rPre, rIn),
	}
	return r
}
