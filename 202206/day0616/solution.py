from typing import List
from collections import Counter


class Solution:
    # 532. 数组中的 k-diff 数对
    def findPairs(self, nums: List[int], k: int) -> int:
        nums_counter = Counter(nums)
        if k == 0:
            return sum(map(lambda e: 1, filter(lambda e: e > 1, nums_counter.values())))
        nums = nums_counter.keys()
        return sum([1 for v in nums if v + k in nums])
