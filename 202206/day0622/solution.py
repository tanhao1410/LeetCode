from typing import Optional, List
from collections import deque


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    # 面试题 17.19. 消失的两个数字
    def missingTwo(self, nums: List[int]) -> List[int]:
        n = len(nums)
        num_sum = sum(nums)
        all_sum = (n + 2) * (n + 3) // 2
        # 即 a + b == all_sum - num_sum
        # 理论上来说，a ,b 应该一个大于 （a + b)//2 ,有一个小于等于，找到这个小于等于的即可
        two_num_sum = (all_sum - num_sum) // 2
        num_sum2 = sum(filter(lambda e: e <= two_num_sum, nums))
        all_sum2 = two_num_sum * (two_num_sum + 1) // 2
        one = all_sum2 - num_sum2
        return [one, two_num_sum - one]

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
