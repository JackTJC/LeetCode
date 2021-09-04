package dfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, target int) (ans [][]int) {
	path := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(n *TreeNode, l int) {
		if n == nil {
			return
		}
		path = append(path, n.Val)
		l -= n.Val
		if n.Left == nil && n.Right == nil && l == 0 {
			ans = append(ans, path[:])
		}
		if n.Left != nil {
			dfs(n.Left, l)
		}
		if n.Right != nil {
			dfs(n.Right, l)
		}
		path = path[:len(path)-1]
	}
	dfs(root, target)
	return
}
