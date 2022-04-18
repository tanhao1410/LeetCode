from typing import List


class TreeNode:
    pass


class Solution:
    #199. 二叉树的右视图
    def rightSideView(self, root: TreeNode) -> List[int]:
        # 层次遍历
        q, res = [], []
        if root is not None:
            q.append(root)
        while len(q) > 0:
            # 拿出队列中的最后一个
            qLen = len(q)
            res.append(q[qLen - 1])
            for i in range(qLen):
                first = q[i]
                if first.left is not None:
                    q.append(first.left)
                if first.right is not None:
                    q.append(first.right)
            q = q[qLen:]
        return res

    # 386. 字典序排数
    def lexicalOrder(self, n: int) -> List[int]:
        res = []
        for i in range(1, 10):
            self.dfs(n, i, res)
        return res

    def dfs(self, limit, next, res: List[int]):
        if next <= limit:
            res.append(next)
            for i in range(10):
                self.dfs(limit, next * 10 + i, res)
