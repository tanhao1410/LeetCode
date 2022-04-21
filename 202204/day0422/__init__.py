from typing import List


class Solution:
    # 396. 旋转函数
    def maxRotateFunction(self, nums: List[int]) -> int:
        # f(0) = 易求
        # f(1) = f(0) + sum(nums) - nums[3] - 3 * nums[3]
        # f(2) = f(1) + sum(nums) - nums[2] - 3 * nums[2]
        # f(3) = f(2) + sum(nums) - nums[1] - 3 * nums[1]
        # 所以求最大值就转变为了看谁最大。
        pre = 0
        nums_sum = sum(nums)
        for i, v in enumerate(nums):
            pre += i * v
        next = 0
        res = pre
        for i in range(1, len(nums)):
            next = pre + nums_sum - len(nums) * nums[len(nums) - i]
            res = max(next, res)
            pre = next
        return res
