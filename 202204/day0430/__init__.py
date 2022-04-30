from typing import List


class Solution:
    # 908. 最小差值 I
    def smallestRangeI(self, nums: List[int], k: int) -> int:
        # return max(0,max(nums) - min(nums) - 2 * k)
        max_, min_ = max(nums), min(nums)
        if max_ - min_ > 2 * k:
            return max_ - min_ - 2 * k
        return 0
