from typing import Optional
from collections import deque


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    # 1382. 将二叉搜索树变平衡
    def balanceBST(self, root: TreeNode) -> TreeNode:
        # 重新构造树，从中间开始构造，
        l = list()
        self.midReadTree(root, l)
        return self.sortedListBST(l)

    def midReadTree(self, root: TreeNode, l):
        if root:
            self.midReadTree(root.left, l)
            l.append(root.val)
            self.midReadTree(root.right, l)

    def sortedListBST(self, l) -> TreeNode:
        root = None
        if l:
            root = TreeNode(val=l[len(l) // 2])
            root.left = self.sortedListBST(l[:len(l) // 2])
            root.right = self.sortedListBST(l[len(l) // 2 + 1:])
        return root

    # 1315. 祖父节点值为偶数的节点和
    def sumEvenGrandparent(self, root: TreeNode) -> int:
        # 思路：递归法：一个节点是偶数的话，那么它的所有孙节点应该加入结果，如果不是，答案为其子节点
        res = 0
        if root:
            if root.val % 2 == 0:
                res += self.sumGrandSon(root)
            res += self.sumEvenGrandparent(root.left)
            res += self.sumEvenGrandparent(root.right)
        return res

    def sumGrandSon(self, root: TreeNode) -> int:
        res = 0
        if root.left:
            res += self.sumSon(root.left)
        if root.right:
            res += self.sumSon(root.right)
        return res

    def sumSon(self, root: TreeNode) -> int:
        res = 0
        if root.left:
            res += root.left.val
        if root.right:
            res += root.right.val
        return res

    # 513. 找树左下角的值
    def findBottomLeftValue(self, root: Optional[TreeNode]) -> int:
        queue = deque()
        queue.append(root)
        res = 0
        while queue:
            queue_len = len(queue)
            res = queue[0].val
            for _ in range(queue_len):
                left = queue.popleft()
                if left.left:
                    queue.append(left.left)
                if left.right:
                    queue.append(left.right)
        return res
