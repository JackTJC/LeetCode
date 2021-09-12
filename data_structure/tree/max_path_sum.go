package tree

import "math"

func maxPathSum(root *TreeNode) int {
	max := func(arr ...int) int {
		var ret = int(math.MinInt32)
		for _, num := range arr {
			if num > ret {
				ret = num
			}
		}
		return ret
	}
	ret := int(math.MinInt32)
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lMax := dfs(root.Left)
		rMax := dfs(root.Right)
		ret = max(ret, lMax+rMax+root.Val, lMax+root.Val, rMax+root.Val, root.Val)
		return max(lMax+root.Val, rMax+root.Val, root.Val)
	}
	dfs(root)
	return ret
}
