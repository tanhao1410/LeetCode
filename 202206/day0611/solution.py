class Solution:
    # 926. 将字符串翻转到单调递增
    def minFlipsMonoIncr(self, s: str) -> int:
        pre_one = [0 for _ in range(len(s))]
        post_zero = [0 for _ in range(len(s))]
        for i in range(len(s)):
            if s[i] == '1':
                pre_one[i] = 1
            if s[len(s) - 1 - i] == '0':
                post_zero[len(s) - 1 - i] = 1
        for i in range(1, len(s)):
            pre_one[i] += pre_one[i - 1]
            post_zero[len(s) - 1 - i] += post_zero[len(s) - i]
        return min(map(lambda e: e[0] + e[1], zip(pre_one, post_zero))) - 1

    # 476. 数字的补数
    def findComplement(self, num: int) -> int:
        bits = []
        while num > 0:
            bits.append(1 - (num & 1))
            num >>= 1
        res = 0
        for i in reversed(bits):
            res <<= 1
            res += i
        return res

    # 762. 二进制表示中质数个计算置位
    def countPrimeSetBits(self, left: int, right: int) -> int:
        def one_count(num: int):
            res = 0
            while num > 0:
                res += 1
                num >>= 1
            return res

        res = 0
        prime_num = [2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31]
        for num in range(left, right + 1):
            if one_count(num) in prime_num:
                res += 1
        return res
