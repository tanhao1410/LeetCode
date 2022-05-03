from functools import reduce
from itertools import count
from typing import List


class Solution:
    # 1295. 统计位数为偶数的数字
    def findNumbers(self, nums: List[int]) -> int:
        def bits_num(num):
            res = 0
            while num > 0:
                res += 1
                num //= 10
            return res

        res = 0
        for num in nums:
            if bits_num(num) % 2 == 0:
                res += 1
        return res

    # 1281. 整数的各位积和之差
    def subtractProductAndSum(self, n: int) -> int:
        nums = []
        while n != 0:
            nums.append(n % 10)
            n //= 10
        return reduce(lambda a, b: a * b, nums) - sum(nums)
