from typing import List
from collections import deque


class Solution:
    # 433. 最小基因变化
    def minMutation(self, start: str, end: str, bank: List[str]) -> int:
        def isDiffOne(p, q):
            diff = 0
            for i in range(8):
                if p[i] != q[i]:
                    diff += 1
            return diff == 1

        res = 0
        queue = deque()
        queue.append(start)
        while len(queue) > 0:
            # 从队列中弹出序列
            count = len(queue)
            for _ in range(count):
                s = queue.popleft()
                if s == end:
                    return res
                need_remove = []
                for b in bank:
                    if isDiffOne(s, b):
                        queue.append(b)
                        need_remove.append(b)
                for r in need_remove:
                    bank.remove(r)
            res += 1
        return -1


solution = Solution()
solution.minMutation('AACCGGTT', 'AAACGGTA', ["AACCGGTA", "AACCGCTA", "AAACGGTA"])
