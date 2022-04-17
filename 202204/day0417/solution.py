from typing import List, Optional


class Solution:
    #819. 最常见的单词
    def mostCommonWord(self, paragraph: str, banned: List[str]) -> str:
        word_map = {}
        banned_set = set(banned)
        start, end = 0, 0
        max_count = 0
        res = ''
        while end <= len(paragraph):
            if end < len(paragraph) and paragraph[end].isalnum():
                end += 1
            else:
                word = paragraph[start:end].lower()
                if word not in banned_set and word_map.get(word) is None:
                    word_map[word] = 1
                elif word not in banned_set:
                    word_map[word] += 1
                if word_map.get(word) is not None and word_map[word] > max_count and word != '':
                    res = word
                    max_count = word_map[word]
                end += 1
                start = end
        return res


s = Solution()
print(s.mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", ['hit']))
