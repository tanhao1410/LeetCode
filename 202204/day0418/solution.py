from typing import List


class Solution:
    #386. 字典序排数
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
