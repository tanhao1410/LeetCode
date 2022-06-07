from typing import List


class Solution:
    # 2119. 反转两次的数字
    def isSameAfterReversals(self, num: int) -> bool:
        return num == 0 or num % 10 != 0

    # 875. 爱吃香蕉的珂珂
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        l, r = 1, max(piles)
        m = (l + r) // 2
        while l < r:
            times = sum(map(lambda c: (c + m - 1) // m, piles))
            if times <= h:
                # 时间够，减小速度
                r = m
            else:
                l = m + 1
            m = (r + l) // 2
        return l
