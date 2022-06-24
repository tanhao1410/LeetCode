from typing import Optional, List
from collections import deque


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    # 剑指 Offer II 046. 二叉树的右侧视图
    def rightSideView(self, root: TreeNode) -> List[int]:
        queue = deque()
        res = []
        if root:
            queue.append(root)
            while queue:
                queue_len = len(queue)
                item = 0
                for _ in range(queue_len):
                    head = queue.popleft()
                    if head.left:
                        queue.append(head.left)
                    if head.right:
                        queue.append(head.right)
                    item = head.val
                res.append(item)
        return res

    # 2265. 统计值等于子树平均值的节点数
    def averageOfSubtree(self, root: Optional[TreeNode]) -> int:
        return self.sumCountRes(root)[2]

    def sumCountRes(self, root: TreeNode) -> List[int]:
        res = [0, 0, 0]
        if root:
            left_res = self.sumCountRes(root.left)
            right_res = self.sumCountRes(root.right)
            # 个数等于左子树个数+右子树个数+1
            res[0] = 1 + left_res[0] + right_res[0]
            # 总和
            res[1] = root.val + left_res[1] + right_res[1]
            res[2] = left_res[2] + right_res[2]
            if res[1] // res[0] == root.val:
                res[2] += 1
        return res

    # 1026. 节点与其祖先之间的最大差值
    def maxAncestorDiff(self, root: Optional[TreeNode]) -> int:
        return self.maxAndMinValue(root)[2]

    # 求一个树的最大值，最小值,最大差距，。
    def maxAndMinValue(self, root) -> (int, int, int):
        min_val, max_val, max_diff = root.val, root.val, 0
        if root.left:
            left_min, left_max, left_diff = self.maxAndMinValue(root.left)
            min_val = min(min_val, left_min)
            max_val = max(max_val, left_max)
            max_diff = max(max_diff, left_diff, abs(root.val - left_min), abs(root.val - left_max))
        if root.right:
            right_min, right_max, right_diff = self.maxAndMinValue(root.right)
            min_val = min(min_val, right_min)
            max_val = max(max_val, right_max)
            max_diff = max(max_diff, right_diff, abs(root.val - right_min), abs(root.val - right_max))
        return (min_val, max_val, max_diff)

    # 515. 在每个树行中找最大值
    def largestValues(self, root: Optional[TreeNode]) -> List[int]:
        res = []
        if root:
            q = deque()
            q.append(root)
            while q:
                q_len = len(q)
                max_value = 0x80000000
                for _ in range(q_len):
                    left = q.popleft()
                    max_value = max(max_value, left.val)
                    if left.left:
                        q.append(left.left)
                    if left.right:
                        q.append(left.right)
                res.append(max_value)
        return res
