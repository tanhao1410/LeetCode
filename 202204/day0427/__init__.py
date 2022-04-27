from typing import List


class Solution:
    # 417. 太平洋大西洋水流问题
    def pacificAtlantic(self, heights: List[List[int]]) -> List[List[int]]:

        m, n = len(heights), len(heights[0])

        def ocean(init) -> set[(int, int)]:
            res = set(init)
            dirt = [(1, 0), (0, 1), (-1, 0), (0, -1)]
            init_len = len(init)
            while init_len > 0:
                for i, j in init:
                    for x, y in dirt:
                        if 0 <= i + x < m and 0 <= j + y < n and heights[i + x][j + y] >= \
                                heights[i][j] and (i + x, j + y) not in res:
                            init.append((i + x, j + y))
                            res.add((i + x, j + y))
                init = init[init_len:]
                init_len = len(init)
            return res

        # 先寻找流向太平洋的点。再寻找流向大西洋的，最后寻找交集。
        pacific_init = [(0, x) for x in range(n)] + [(x, 0) for x in range(m)]
        print(pacific_init)
        atalntic_init = [(m - 1, x) for x in range(n)] + [(x, n - 1) for x in range(m)]
        return [list(lo) for lo in ocean(pacific_init).intersection(ocean(atalntic_init))]


solution = Solution()
solution.pacificAtlantic([[1, 1], [1, 1], [1, 1]])
