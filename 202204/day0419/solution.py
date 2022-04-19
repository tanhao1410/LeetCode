from typing import List


class Solution:
    # 821. 字符的最短距离
    def shortestToChar(self, s: str, c: str) -> List[int]:
        res = [-1 for x in range(len(s))]
        # 找到所有c的位置
        local = [x for x in range(len(s)) if s[x] == c]
        for i in local:
            res[i] = 0
        if len(local) == 1:
            return [abs(x - local[0]) for x in range(len(s))]
        one = local[0]
        two = local[1]
        next = 2
        for i in range(len(s)):
            if res[i] == -1:
                res[i] = min(abs(one - i), abs(two - i))
            if i > two and next < len(local):
                one, two = two, local[next]
                next += 1
        return res
