package tree

import "testing"

func TestBuildTree(t *testing.T) {
	buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
}

func TestIsSub(t *testing.T) {
	isSubStructure(&TreeNode{
		Val:   -2,
		Left:  &TreeNode{1, nil, nil},
		Right: &TreeNode{-1, nil, nil},
	}, &TreeNode{
		Val:   -2,
		Left:  &TreeNode{1, nil, nil},
		Right: &TreeNode{1, nil, nil},
	})
}
