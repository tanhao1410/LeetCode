from typing import List


class Solution:
    # 7. 整数反转
    def reverse(self, x: int) -> int:
        # 计算每一位数字，共32位，如何判断是否大于最大的数呢，或小于最小的数呢？
        # 最大的正数 2**31 - 1
        bits = []
        flag = x > 0
        x = abs(x)
        while x > 0:
            bits.append(x % 10)
            x //= 10
        # 需要判断是否越界
        max_int = 2 ** 31 - 1
        res = 0
        for bit in bits:
            res *= 10
            res += bit
        if flag:
            if res > max_int:
                return 0
            else:
                return res
        else:
            if res > max_int + 1:
                return 0
            else:
                return -res

    # 64. 最小路径和
    def minPathSum(self, grid: List[List[int]]) -> int:
        for i in range(len(grid)):
            for j in range(len(grid[0])):
                if i == 0:
                    if j > 0:
                        grid[i][j] += grid[i][j - 1]
                elif j == 0:
                    grid[i][j] += grid[i - 1][j]
                else:
                    grid[i][j] += min(grid[i - 1][j], grid[i][j - 1])
        return grid[-1][-1]

    # 5. 最长回文子串
    def longestPalindrome(self, s: str) -> str:
        # 动态规划，dp[i][j] 如果s[i] == s[j] 则看dp[i + 1][j - 1]是否是回文。
        # 两层循环，
        dp = [[False for _ in range(len(s))] for _ in range(len(s))]
        max, res = 0, ''
        for i in range(len(s) - 1, -1, -1):
            for j in range(i, len(s)):
                if s[i] == s[j]:
                    if i == j or i + 1 == j:
                        dp[i][j] = True
                    else:
                        dp[i][j] = dp[i + 1][j - 1]
                if dp[i][j]:
                    if j - i + 1 > max:
                        max = j - i + 1
                        res = s[i:j + 1]
        return res

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


s = Solution()
s.longestPalindrome("babad")
