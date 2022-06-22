from typing import Optional
from collections import deque


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
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
