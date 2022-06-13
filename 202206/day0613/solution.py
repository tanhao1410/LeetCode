from typing import List


class Solution:
    # 1051. 高度检查器
    def heightChecker(self, heights: List[int]) -> int:
        return sum(map(lambda e: 1, filter(lambda e: e[0] != e[1], zip(sorted(heights), heights))))
