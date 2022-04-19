class Solution:
    # 415. 字符串相加
    def addStrings(self, num1: str, num2: str) -> str:
        # 先对齐，然后从后面往前面加，注意进位即可
        i = 0
        flag = 0
        res = ''
        while i < len(num1) or i < len(num2):
            num = flag
            if i < len(num1):
                num += int(num1[len(num1) - 1 - i])
            if i < len(num2):
                num += int(num2[-i - 1])
            flag = num // 10
            res = str(num % 10) + res
            i += 1
        if flag != 0:
            res = '1' + res
        return res
