# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def isValidBST(self, root: TreeNode) -> bool:
        def subValid(root, minVal, maxVal):
            if root.val <= minVal or root.val >= maxVal:
                return False
            rightValid = True
            leftValid = True
            if root.left:
                leftValid = subValid(root.left, minVal, min(root.val, maxVal))
            if root.right:
                rightValid = subValid(root.right, max(minVal, root.val), maxVal)
            return (rightValid and leftValid)

        if root is None:
            return True
        else:
            return subValid(root, -2 ** 31 - 1, 2 ** 31)
