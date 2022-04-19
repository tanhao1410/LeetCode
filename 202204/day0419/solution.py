from typing import List, Optional


class TreeNode:
    pass


class Solution:
    # 124. 二叉树中的最大路径和
    def maxPathSum(self, root: Optional[TreeNode]) -> int:
        # 思路：对于任意一个节点来说，以自己为末尾的链条。以自己为中心的链条。以自己为开始的链条。三种
        # 任意一个节点，所能形成的最大和为三者之一。
        # 怎么求以自己为开始的最大的呢。this.val + this.right(最大开始或最大结束或0)
        # 以自己为结束的最大值：this.val + this.left(最大开始或最大结束或0)
        # 以自己为中心的最大值：this.val + this.left(最大开始或最大结束或0) + this.right(最大开始或最大结束或0)
        return self.maxStartOrEnd(root)[2]

    # 返回值分别为最大结束，最大开始，最大值
    def maxStartOrEnd(self, root: Optional[TreeNode]) -> (int, int, int):
        if root is None:
            return (-10000, -10000, -100000)
        # 先求最大结束
        leftEnd, leftStart, leftMax = self.maxStartOrEnd(root.left)
        rightEnd, rightStart, rightMax = self.maxStartOrEnd(root.right)
        curEnd = root.val + max(0, leftStart, leftEnd)
        curStart = root.val + max(0, rightStart, rightEnd)
        curMax = root.val + max(0, leftStart, leftEnd) + max(0, rightEnd, rightStart)
        return (curEnd, curStart, max(curMax, leftMax, rightMax))

    # 821. 字符的最短距离
    def shortestToChar(self, s: str, c: str) -> List[int]:
        res = [-1 for x in range(len(s))]
        # 找到所有c的位置
        local = [x for x in range(len(s)) if s[x] == c]
        for i in local:
            res[i] = 0
        if len(local) == 1:
            return [abs(x - local[0]) for x in range(len(s))]
        one = local[0]
        two = local[1]
        next = 2
        for i in range(len(s)):
            if res[i] == -1:
                res[i] = min(abs(one - i), abs(two - i))
            if i > two and next < len(local):
                one, two = two, local[next]
                next += 1
        return res
