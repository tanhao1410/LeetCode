from typing import Optional, List
from collections import deque


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
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
