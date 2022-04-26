from typing import List


class Solution:
    # 883. 三维形体投影面积
    def projectionArea(self, grid: List[List[int]]) -> int:
        # x视图
        res = sum([sum([1 for num in v if num > 0]) for v in grid])
        # for inner in grid:
        #     res += max(inner)
        res += sum([max(v) for v in grid])
        # for i in range(len(grid[0])):
        #     res += max([grid[v][i] for v in range(len(grid))])
        res += sum([max([grid[x][y] for x in range(len(grid))]) for y in range(len(grid[0]))])
        return res
