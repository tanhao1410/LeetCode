from typing import List


class Node:
    pass


class Solution:
    # 427. 建立四叉树
    def construct(self, grid: List[List[int]]) -> 'Node':
        first = grid[0][0]
        grid_len = len(grid)
        if len(grid) == 1:
            return Node(first, True, None, None, None, None)
        # 确认所有的是否是一个值
        for row in grid:
            for n in row:
                if first != n:
                    # 分割成四份
                    topLeft = self.construct([row[:grid_len // 2] for row in grid[:grid_len // 2]])
                    topRight = self.construct([row[grid_len // 2:] for row in grid[:grid_len // 2]])
                    bottomLeft = self.construct([row[:grid_len // 2] for row in grid[grid_len // 2:]])
                    bottomRight = self.construct([row[grid_len // 2:] for row in grid[grid_len // 2:]])
                    return Node(1, False, topLeft, topRight, bottomLeft, bottomRight)
        return Node(first, True, None, None, None, None)
