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
