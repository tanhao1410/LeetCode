from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    # 1022. 从根到叶的二进制数之和
    def sumRootToLeaf(self, root: Optional[TreeNode], pre_value=0) -> int:
        pre_value = pre_value * 2 + root.val
        if root.left is None and root.right is None:
            return pre_value
        res = 0
        if root.left:
            res = self.sumRootToLeaf(root.left, pre_value)
        if root.right:
            res += self.sumRootToLeaf(root.right, pre_value)
        return res
