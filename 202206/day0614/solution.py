from typing import List, Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    # 1123. 最深叶节点的最近公共祖先
    def lcaDeepestLeaves(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        # 层级遍历法
        rows = [[root]]
        while True:
            # 取上面一层
            up_row = rows[-1]
            new_row = []
            for node in up_row:
                if node.left:
                    node.left.parent = node
                    new_row.append(node.left)
                if node.right:
                    node.right.parent = node
                    new_row.append(node.right)
            if new_row:
                rows.append(new_row)
            else:
                break
        # 从最下层开始，如果最下层只有一个节点，直接返回即可。多个节点的话，找共同祖先
        row = set(rows[-1])
        parents = set()
        while len(row) > 1:
            parents.clear()
            for node in row:
                parents.add(node.parent)
            row.clear()
            row = row.union(parents)
        return row.pop()

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
