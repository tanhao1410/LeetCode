from functools import reduce
from typing import List


class Solution:
    # 1281. 整数的各位积和之差
    def subtractProductAndSum(self, n: int) -> int:
        nums = []
        while n != 0:
            nums.append(n % 10)
            n //= 10
        return reduce(lambda a, b: a * b, nums) - sum(nums)
