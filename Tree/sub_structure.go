package tree

/*
剑指 Offer 26. 树的子结构
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
B是A的子结构， 即 A中有出现和B相同的结构和节点值。
*/
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil && B != nil || A != nil && B == nil {
		return false
	}
	if A == nil && B == nil {
		return true
	}
	if A.Val == B.Val && compare(A, B) {
		return true
	}
	l, r := isSubStructure(A.Left, B), isSubStructure(A.Right, B)
	return l || r
}

func compare(subA *TreeNode, subB *TreeNode) bool {
	if subA == nil && subB != nil || subA != nil && subB == nil {
		return false
	}
	if subA == nil && subB == nil {
		return true
	}
	if subA.Val != subB.Val {
		return false
	}
	left, right := true, true
	if subB.Left != nil {
		left = compare(subA.Left, subB.Left)
	}
	if subB.Right != nil {
		right = compare(subA.Right, subB.Right)
	}
	return left && right
}
