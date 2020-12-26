# 1080
# DFS
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def sufficientSubset(self, root: TreeNode, limit: int) -> TreeNode:
        def deleteNode(root, sum):
            if root.left is None and root.right is None:
                if root.val + sum < limit:
                    return True
                else:
                    return False
            sum += root.val
            if root.left:
                if deleteNode(root.left, sum):
                    root.left = None
            if root.right:
                if deleteNode(root.right, sum):
                    root.right = None
            if root.left is None and root.right is None:
                return True
            return False

        if root is None:
            return None
        returnVal = deleteNode(root, 0)
        if returnVal:
            return None
        return root
