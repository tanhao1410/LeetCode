from typing import List

class NestedInteger:
    pass

class Solution:
    # 385.迷你语法分析器
    def deserialize(self, s: str) -> NestedInteger:
        # 思路：如果是单独的数字，直接生成一个返回即可。
        # 如果不是数字，说明是list，将list里的每一个元素按照递归形式生成对象，放入里面即可。
        res = NestedInteger()
        if s == '[]':
            return res
        elif s.startswith('['):
            list = self.process(s)
            for ele in list:
                res.add(self.deserialize(ele))
        else:
            res.setInteger(int(s))
        return res

    # 将[1,2,[2,3]]拆分成list，
    def process(self, s: str) -> List:
        # 先把左右两边的括号去除掉
        res = []
        # 思路：用一个变量表示前面有多少【,如果前面【是零个，一旦读到,或读到结尾则生成一个新的
        # 否则，继续往下读。
        flag = 0
        start = 1
        end = 1
        while end < len(s) - 1:
            # 如果开始是[，说明内部元素是列表
            if s[end] == '[':
                flag += 1
            elif s[end] == ']':
                flag -= 1
            elif s[end] == ',' and flag == 0:
                # 此时需要生成一个元素
                res.append(s[start:end])
                start = end + 1
            end += 1
        # 最后一个元素加入进来
        res.append(s[start:end])
        return res
