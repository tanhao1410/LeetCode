from typing import List


class Solution:
    # 921. 使括号有效的最少添加
    def minAddToMakeValid(self, s: str) -> int:
        # 思路：从左往右，依次添加括号，对于已经完成匹配的，删除掉。如果前面是)，则添加对应的（
        # 如果对应的是（，则继续往前走，并记录)的数量。如果一旦达到平衡，则去掉，如果达不到平衡，则在最后的时候补充上去
        res = 0
        left = 0
        for w in s:
            if w == '(':
                left += 1
            elif left > 0:
                left -= 1
            else:
                res += 1
        return res + left

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
