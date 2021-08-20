package tree

type treeBuilder struct {
	preorder, inorder []int
}

//四个参数分别代表列表范围
func (tb *treeBuilder) run(pl, pr, il, ir int) *TreeNode {
	if pl == pr {
		return nil
	}
	var cutIdx int
	for i := il; i < ir; i++ {
		if tb.inorder[i] == tb.preorder[pl] {
			cutIdx = i
		}
	}
	ret := &TreeNode{
		Val: tb.preorder[pl],
	}
	lil, lir, ril, rir = il, il+cutIdx, il+cutIdx+1, ir
	llp
	ret.Left = tb.run()
	ret.Right = tb.run()
	return ret
}

func buildTreeOpt(preorder []int, inorder []int) *TreeNode {
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
