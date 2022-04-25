import random
from typing import List


# 398. 随机数索引
class Solution:

    def __init__(self, nums: List[int]):
        map = {}
        for i, num in enumerate(nums):
            if num in map:
                map[num].append(i)
            else:
                map[num] = [i]
        self.map = map

    def pick(self, target: int) -> int:
        return random.choice(self.map[target])
