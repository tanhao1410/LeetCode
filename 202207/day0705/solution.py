class Solution:
    # 2169. 得到 0 的操作数
    def countOperations(self, num1: int, num2: int) -> int:
        res = 0
        while num1 != 0 and num2 != 0:
            res += 1
            if num1 >= num2:
                num1 = num1 - num2
            else:
                num2 = num2 - num1
        return res
