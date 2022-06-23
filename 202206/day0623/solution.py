from collections import Counter
from typing import List


class Solution:
    # 30. 串联所有单词的子串
    def findSubstring(self, s: str, words: List[str]) -> List[int]:
        word_count = Counter(words)
        word_len = len(words[0])
        # 窗口的大小，单词的数量
        window = len(words)
        res = []
        for i in range(word_len):
            # i即窗口开始的位置。加入len(words)个单词，
            start = i
            window_count = Counter()
            for j in range(window):
                window_count[s[start + j * word_len:start + (j + 1) * word_len]] += 1
            if word_count == window_count:
                res.append(i)

            # 开始移动窗口，j代表窗口的尾部，每次移动一个单词的长度，直到s结尾
            for j in range(i + word_len * (window + 1), len(s) + 1, word_len):
                # 减去前面一个单词
                window_count[s[start:start + word_len]] -= 1
                # 开始位置移动了一个单词
                start += word_len
                # 加上新加进来的单词
                window_count[s[j - word_len:j]] += 1
                if word_count == window_count:
                    res.append(start)
        return res

    # 1387. 将整数按权重排序
    def getKth(self, lo: int, hi: int, k: int) -> int:
        map_count = {1 << i: i for i in range(32)}
        return sorted(range(lo, hi + 1), key=lambda i: self.get_count(i, map_count))[k - 1]

    def get_count(self, num, map_count) -> int:
        if num in map_count:
            return map_count[num]
        if num % 2 == 0:
            res = self.get_count(num // 2, map_count)
        else:
            res = self.get_count(num * 3 + 1, map_count)
        map_count[num] = res + 1
        return map_count[num]


# 1146. 快照数组
class SnapshotArray:

    def __init__(self, length: int):
        self.datas = {}
        self.snap_num = 0

    def set(self, index: int, val: int) -> None:
        values = self.datas.get(index, [])
        if len(values) > 0 and values[-1][0] == self.snap_num:
            values[-1] = (self.snap_num, val)
        else:
            values.append((self.snap_num, val))
        self.datas[index] = values

    def snap(self) -> int:
        self.snap_num += 1
        return self.snap_num - 1

    def get(self, index: int, snap_id: int) -> int:
        if index in self.datas:
            values = self.datas[index]
            for i in range(len(values) - 1, -1, -1):
                if values[i][0] <= snap_id:
                    return values[i][1]
        return 0
