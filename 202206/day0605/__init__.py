from typing import List
from random import random


class Solution:
    # 1528. 重新排列字符串
    def restoreString(self, s: str, indices: List[int]) -> str:
        res = ''
        for i in sorted(zip(s, indices), key=lambda e: e[1]):
            res += i[0]
        return res

    # 1417. 重新格式化字符串
    def reformat(self, s: str) -> str:
        letters, nums = [], []
        for i in s:
            if '0' <= i <= '9':
                nums.append(i)
            else:
                letters.append(i)
        if abs(len(letters) - len(nums)) > 1:
            return ''
        res = ''
        for i in range(min(len(letters), len(nums))):
            res += letters[i] + nums[i]
        if len(letters) > len(nums):
            res += letters[-1]
        elif len(letters) < len(nums):
            res = nums[-1] + res
        return res

    # 375. 猜数字大小 II
    def getMoneyAmount(self, n: int) -> int:
        dp = [[0 for _ in range(n + 1)] for _ in range(n + 1)]
        for start in reversed(range(n + 1)):
            for end in range(start + 1, n + 1):
                max_spend = 100000
                for i in range(start, end + 1):
                    cur_spend = 0
                    if i > start:
                        cur_spend = max(cur_spend, dp[start][i - 1] + i)
                    if i < end:
                        cur_spend = max(cur_spend, dp[i + 1][end] + i)
                    max_spend = min(max_spend, cur_spend)
                dp[start][end] = max_spend
        return dp[0][n]

    # 478. 在圆内随机生成点
    def __init__(self, radius: float, x_center: float, y_center: float):
        self.radius = radius
        self.x = x_center
        self.y = y_center

    def pointInCircle(self, x, y) -> bool:
        return self.radius ** 2 >= (x - self.x) ** 2 + (y - self.y) ** 2

    def randPoint(self) -> List[float]:
        # 看是否落在了园内
        zero_point = [self.x - self.radius, self.y - self.radius]
        rand_point = [random() * 2 * self.radius + zero_point[0], random() * 2 * self.radius + zero_point[1]]
        if self.pointInCircle(rand_point[0], rand_point[1]):
            return rand_point
        return self.randPoint()
