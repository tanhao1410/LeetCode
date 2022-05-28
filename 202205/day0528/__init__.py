class Solution:
    # 1021. 删除最外层的括号
    def removeOuterParentheses(self, s: str) -> str:
        # 思路：通过一个变量来记录左括号的数量，一旦左括号为0的话，说明前面切出来的为一个 符合逻辑的括号，将它交由单个括号处理函数处理
        # 单个函数处理方法，若长度为2，直接返回，若长度不为2，则脱去外面的括号，里面的 递归处理即可
        start_index, left_count = 0, 0
        res = ''
        for i in range(len(s)):
            if s[i] == '(':
                left_count += 1
            else:
                left_count -= 1
            if left_count == 0:
                res += s[start_index + 1:i]
                start_index = i + 1
        return res
