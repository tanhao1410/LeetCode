from typing import List
from random import random


class Solution:
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
