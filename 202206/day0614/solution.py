from typing import List


class Solution:
    # 498. 对角线遍历
    def findDiagonalOrder(self, mat: List[List[int]]) -> List[int]:
        res = []
        m, n = len(mat), len(mat[0])
        up = True
        start = [0, 0]
        while len(res) < m * n:
            res.append(mat[start[0]][start[1]])
            if up:
                if start[0] - 1 >= 0 and start[1] + 1 < n:
                    start[0] -= 1
                    start[1] += 1
                elif start[1] + 1 < n:
                    # 往右走
                    start[1] += 1
                    up = not up
                else:
                    start[0] += 1
                    up = not up
            else:
                if start[0] + 1 < m and start[1] - 1 >= 0:
                    start[0] += 1
                    start[1] -= 1
                elif start[0] + 1 < m:
                    # 往下走
                    start[0] += 1
                    up = not up
                else:
                    start[1] += 1
                    up = not up
        return res
