from typing import Optional, List
from collections import deque


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
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
