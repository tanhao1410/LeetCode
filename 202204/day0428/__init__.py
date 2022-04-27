from typing import List


class Solution:
    # 905. 按奇偶排序数组
    def sortArrayByParity(self, nums: List[int]) -> List[int]:
        # list(filter(lambda num : num % 2 == 0,nums)) + list(filter(lambda  num : num % 2,nums))
        even, odd = 0, len(nums) - 1
        while even < odd:
            while even < len(nums) and nums[even] % 2 == 0:
                even += 1
            while odd >= 0 and nums[odd] % 2 == 1:
                odd -= 1
            if even < odd:
                nums[even], nums[odd] = nums[odd], nums[even]
        return nums
