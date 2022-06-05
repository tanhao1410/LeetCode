from typing import List
from random import random


class Solution:
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


s = Solution(1.0, 2.0, 3.0)
print(s.select_num(1, 25))
