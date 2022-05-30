from typing import Optional, List


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    # 682. 棒球比赛
    def calPoints(self, ops: List[str]) -> int:
        scores = []
        for x in ops:
            if x == '+':
                scores.append(scores[-1] + scores[-2])
            elif x == 'D':
                scores.append(scores[-1] * 2)
            elif x == 'C':
                scores = scores[:-1]
            else:
                scores.append(int(x))
        return sum(scores)

    # 655. 输出二叉树
    def printTree(self, root: TreeNode, height=0) -> List[List[str]]:
        m = max(self.treeHeight(root), height)
        if m == 0:
            return [[]]
        n = 2 ** m - 1
        # 递归式解决方案，对于它的左边来说，对于它的右边来说。
        # 只关心自己这一层而已。下一层怎么解决呢？
        # 如果自己本身是空怎么办？
        res = [['' for _ in range(n)]]
        if root:
            res[0][n // 2] = str(root.val)
            left_res = self.printTree(root.left, m - 1)
            right_res = self.printTree(root.right, m - 1)
        else:
            left_res = self.printTree(None, m - 1)
            right_res = self.printTree(None, m - 1)
        for i in range(1, m):
            if left_res[i - 1]:
                left_res[i - 1].append('')
                res.append(left_res[i - 1] + right_res[i - 1])
        return res

    def treeHeight(self, root: TreeNode) -> int:
        if root:
            return max(self.treeHeight(root.left), self.treeHeight(root.right)) + 1
        return 0

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
