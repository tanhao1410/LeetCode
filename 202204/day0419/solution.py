from typing import List, Optional


class TreeNode:
    pass


class Solution:
    # 103. 二叉树的锯齿形层序遍历
    def zigzagLevelOrder(self, root: TreeNode) -> List[List[int]]:
        queue = []
        res = []
        flag = True
        if root:
            queue.append(root)
        while queue:
            length = len(queue)
            item = [0 for _ in range(length)]
            for i in range(length):
                if queue[i].left:
                    queue.append(queue[i].left)
                if queue[i].right:
                    queue.append(queue[i].right)
                if flag:
                    item[i] = queue[i].val
                else:
                    item[length - 1 - i] = queue[i].val
            queue = queue[length:]
            flag = not flag
            res.append(item)
        return res

    # 53. 最大子数组和
    def maxSubArray(self, nums: List[int]) -> int:
        # 思路：
        dp = [nums[0]]
        for i in range(1, len(nums)):
            if dp[-1] > 0:
                dp.append(dp[-1] + nums[i])
            else:
                dp.append(nums[i])
        return max(dp)

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
