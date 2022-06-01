from typing import List


class Solution:
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
