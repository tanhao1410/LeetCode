from typing import List, Optional


class Node:
    pass


class TreeNode:
    pass


class Solution:
    #951. 翻转等价二叉树
    def flipEquiv(self, root1: Optional[TreeNode], root2: Optional[TreeNode]) -> bool:
        # 思路：如果根相同，则继续判断。根不同，则返回false
        # 判断左右子树是否相等。如果不相等，看交换后是否相等。交换后也不相等，返回false
        # 所有的子都相等后，才返回true
        if root1 and not root2 or not root1 and root2:
            return False
        if not root1 and not root2:
            return True
        if root1.val == root2.val:
            # 判断它的左右子树是否满足
            return self.flipEquiv(root1.left, root2.left) and self.flipEquiv(root1.right, root2.right) \
                   or self.flipEquiv(root1.right, root2.left) and self.flipEquiv(root1.left, root2.right)
        return False

    # 661. 图片平滑器
    def imageSmoother(self, img: List[List[int]]) -> List[List[int]]:
        res = [row.copy() for row in img]

        def image_average(x, y) -> int:
            sum_, count = 0, 0
            for i in range(-1, 2):
                for j in range(-1, 2):
                    if 0 <= x + i < len(img) and 0 <= y + j < len(img[0]):
                        sum_ += img[x + i][y + j]
                        count += 1
            return sum_ // count

        for i in range(len(img)):
            for j in range(len(img[0])):
                res[i][j] = image_average(i, j)
        return res

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
