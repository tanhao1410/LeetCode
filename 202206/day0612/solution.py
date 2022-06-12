from typing import List


class Solution:
    # 890. 查找和替换模式
    def findAndReplacePattern(self, words: List[str], pattern: str) -> List[str]:
        res = []
        for word in words:
            # 判断是否符合
            match_dict = dict(zip(word, pattern))
            match_dict2 = dict(zip(pattern, word))
            if all(map(lambda e: match_dict[e[0]] == e[1] and match_dict2[e[1]] == e[0], zip(word, pattern))):
                res.append(word)
        return res
