from typing import List


class Solution:
    # 241. 为运算表达式设计优先级
    def diffWaysToCompute(self, expression: str) -> List[int]:
        # 第一个数字，可以支持到哪个操作符。
        exprees = ['+', '-', '*']
        res = []
        for i in range(len(expression)):
            if expression[i] in exprees:
                # 遇到了一个操作符了，从这开始截断
                prev = self.diffWaysToCompute(expression[:i])
                pro = self.diffWaysToCompute(expression[i + 1:])
                for a in prev:
                    for b in pro:
                        res.append(self.oprator(expression[i])(a, b))
        if not res:
            res.append(int(expression))
        return res

    def oprator(self, oper: str):
        def sum(a, b):
            return a + b

        def mul(a, b):
            return a * b

        def sub(a, b):
            return a - b

        if oper == '+':
            return sum
        if oper == '*':
            return mul
        return sub
