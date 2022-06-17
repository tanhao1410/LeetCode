from collections import Counter
from typing import List


class Solution:
    # 1577 数的平方等于两数乘积的方法数
    def numTriplets(self, nums1: List[int], nums2: List[int]) -> int:
        return self.numTriplets2(nums2, nums1) + self.numTriplets2(nums1, nums2)

    def numTriplets2(self, nums1: List[int], nums2: List[int]) -> int:
        nums2_counter = Counter(nums2)
        res = 0
        for num in nums1:
            num = num * num
            for num2 in nums2_counter.keys():
                if num % num2 == 0 and num / num2 in nums2_counter.keys():
                    if num2 == num / num2:
                        n = nums2_counter[num2]
                        res += n * (n - 1) // 2
                    elif num2 > num / num2:
                        res += nums2_counter[num2] * nums2_counter[num / num2]
        return res
