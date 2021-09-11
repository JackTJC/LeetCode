package backtrack_search

import "testing"

func TestHelloWorld(t *testing.T) {
	movingCount(1, 2, 1)
}

func TestPathSum(t *testing.T) {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{11, &TreeNode{7, nil, nil}, &TreeNode{2, nil, nil}},
			Right: &TreeNode{},
		},
		Right: &TreeNode{
			Val:   8,
			Left:  &TreeNode{13, nil, nil},
			Right: &TreeNode{4, &TreeNode{5, nil, nil}, &TreeNode{1, nil, nil}},
		},
	}
	pathSum(tree, 22)
}
