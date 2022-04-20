class Solution:
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
