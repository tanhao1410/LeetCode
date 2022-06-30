class Solution:
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
