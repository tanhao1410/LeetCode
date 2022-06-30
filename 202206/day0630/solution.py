from typing import List


class Solution:
    # 2215. 找出两数组的不同
    def findDifference(self, nums1: List[int], nums2: List[int]) -> List[List[int]]:
        nums1, nums2 = set(nums1), set(nums2)
        return [[n for n in nums1 if n not in nums2], [n for n in nums2 if n not in nums1]]

    # 1175. 质数排列
    def numPrimeArrangements(self, n: int) -> int:
        MOD = 1000000007
        prime_count = sum([1 for e in range(1, n + 1) if self.isPrime(e)])
        res = 1
        for i in range(2, prime_count + 1):
            res = res * i % MOD
        for i in range(2, n - prime_count + 1):
            res = res * i % MOD
        return res

    def isPrime(self, n: int) -> bool:
        for i in range(2, n):
            if n % i == 0:
                return False
        return True and n > 1
