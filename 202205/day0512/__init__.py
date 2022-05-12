from typing import List


class Solution:
    # 806. 写字符串需要的行数
    def numberOfLines(self, widths: List[int], s: str) -> List[int]:
        row, rem = 1, 0
        for i in s:
            # 确认i占多少位置
            width = widths[ord(i) - ord('a')]
            # 剩余多少宽度 100 - rem
            if width > 100 - rem:
                row += 1
                rem = width
            else:
                rem += width
        return [row, rem]

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
