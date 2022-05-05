from typing import List


class Solution:
    # 713. 乘积小于 K 的子数组
    def numSubarrayProductLessThanK(self, nums: List[int], k: int) -> int:
        left, right = 0, 0
        sum = nums[0]
        res = 0
        while right < len(nums):
            if sum < k:
                res += (right - left + 1)
                right += 1
                if right < len(nums):
                    sum *= nums[right]
            else:
                sum //= nums[left]
                left += 1
                if left > right:
                    right = left
                    if left < len(nums):
                        sum = nums[left]
        return res
