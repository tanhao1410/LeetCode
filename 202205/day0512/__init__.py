from typing import List


class Solution:
    # 944. 删列造序
    def minDeletionSize(self, strs: List[str]) -> int:
        res = 0
        m, n = len(strs), len(strs[0])
        for col in range(n):
            for row in range(1, m):
                if strs[row][col] < strs[row - 1][col]:
                    res += 1
                    break
        return res
