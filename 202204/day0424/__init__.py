class Solution:
    # 868. 二进制间距
    def binaryGap(self, n: int) -> int:
        # 先求一个数的二进制形式
        bits = []
        while n > 0:
            bits.append(n % 2)
            n //= 2
        # 求最大间距
        res, pre_one = 0, 0
        for i, v in enumerate(reversed(bits)):
            if i == v:
                res, pre_one = max(res, i - pre_one), i
        return res
