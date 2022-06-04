from typing import List


class Solution:
    # 779. 第K个语法符号
    def kthGrammar(self, n: int, k: int) -> int:
        if n == 1:
            return 0
        if k <= 2 ** (n - 2):
            return self.kthGrammar(n - 1, k)
        # n是偶数式，不对称，是反过来的
        if n % 2 == 1:
            return self.kthGrammar(n - 1, 2 ** (n - 1) - k + 1)
        return self.kthGrammar(n - 1, 2 ** (n - 1) - k + 1) ^ 1

    # 1544. 整理字符串
    def makeGood(self, s: str) -> str:
        for i in range(len(s) - 1):
            if abs(ord(s[i]) - ord(s[i + 1])) == abs(ord('a') - ord('A')):
                return self.makeGood(s[:i] + s[i + 2:])
        return s

    # 1572. 矩阵对角线元素的和
    def diagonalSum(self, mat: List[List[int]]) -> int:
        res = sum([mat[i][i] + mat[i][len(mat) - 1 - i] for i in range(len(mat))])
        if len(mat) % 2 == 1:
            res -= mat[len(mat) // 2][len(mat) // 2]
        return res

    # 929. 独特的电子邮件地址
    def numUniqueEmails(self, emails: List[str]) -> int:
        email_set = set()
        for email in emails:
            name_host = email.split('@')
            name = name_host[0].replace('.', '').split('+')[0]
            # 如果有.则消除
            email_set.add(name + '@' + name_host[1])
        return len(email_set)
