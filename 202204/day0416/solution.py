from typing import List


class Solution:

    # 1. 两数之和
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        # 采用map的形式，返回的是下标
        map = {}
        for i in range(len(nums)):
            # 看map中是否含有target-nums[i]
            if map.get(target - nums[i]) is not None:
                return [map[target - nums[i]], i]
            map[nums[i]] = i

    # 479. 最大回文数乘积
    def largestPalindrome(self, n: int) -> int:
        if n == 1:
            return 9
        # 两个n位数相乘的话，得到的结果应该是 2*n - 1 或 2n位
        # 先看2n位的，然后在看少一位的大小
        # 98 89 1001 99 999
        part = 10 ** n - 1
        for i in range(part, part // 10, -1):
            # 得到回文数
            num = i
            while i != 0:
                # 得到part的最低位，
                low = i % 10
                i //= 10
                num *= 10
                num += low
            # 寻找合适的数
            j = 10 ** n - 1
            while j * j >= num:
                if num % j == 0:
                    return num % 1337
                j -= 1


for i in range(1, 9):
    s = Solution()
    print(s.largestPalindrome(i))
