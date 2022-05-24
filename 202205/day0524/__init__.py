# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    # 965. 单值二叉树
    def isUnivalTree(self, root: TreeNode) -> bool:
        # 如果有子节点，若子节点不是单值，返回false
        res = True
        if root.left:
            res = root.left.val == root.val
            res = res and self.isUnivalTree(root.left)
        if root.right:
            res = res and root.right.val == root.val
            res = res and self.isUnivalTree(root.right)
        return res
