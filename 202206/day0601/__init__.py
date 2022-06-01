from typing import List


class Solution:
    # 953. 验证外星语词典
    def isAlienSorted(self, words: List[str], order: str) -> bool:
        order2 = 'abcdefghijklmnopqrstuvwxyz'
        letter_dict = dict(zip(order, order2))

        def change_word(word):
            res = ''
            for l in word:
                res += letter_dict.get(l)
            return res

        not_sorted = list(map(change_word, words))
        return sorted(not_sorted) == not_sorted

    # 1071. 字符串的最大公因子
    def gcdOfStrings(self, str1: str, str2: str) -> str:
        # 思路：如果是子串，直接返回一个结果。否则，从短的哪个中寻找，可以是循环一次，可以是循环两次，可以是循环n次。
        for i in range(1, len(str1) + 1):
            if len(str1) % i == 0 and self.isChildStr(str1, str1[:len(str1) // i]):
                if self.isChildStr(str2, str1[:len(str1) // i]):
                    return str1[:len(str1) // i]
        return ''

    def isChildStr(self, parent: str, child: str) -> bool:
        if len(parent) % len(child) != 0:
            return False
        for i in range(len(parent) // len(child)):
            if parent[i * len(child):(1 + i) * len(child)] != child:
                return False
        return True

    # 473. 火柴拼正方形
    def makesquare(self, matchsticks: List[int], edges=None, cur=0) -> bool:
        if cur == len(matchsticks):
            return all(map(lambda e: e == edges[0], edges))
        all_sum = sum(matchsticks)
        if not edges:
            edges = [0, 0, 0, 0]
        if all_sum % 4 == 0:
            # 4个集合，分别代表四个边，每一次从剩余的火柴中选择一个，放在其中的一个
            cur_len = matchsticks[cur]
            for i in range(4):
                if edges[i] + cur_len <= all_sum // 4:
                    edges[i] += cur_len
                    inner_res = self.makesquare(matchsticks, edges, cur + 1)
                    if inner_res:
                        return True
                    # 恢复现场
                    edges[i] -= cur_len
        return False
