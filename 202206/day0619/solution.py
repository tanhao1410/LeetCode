from typing import List


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    # 508. 出现次数最多的子树元素和
    def findFrequentTreeSum(self, root: TreeNode) -> List[int]:
        count_map = {}
        self.treeSum(root, count_map)
        max_count = 0
        for _, c in count_map.items():
            max_count = max(max_count, c)
        return [e[0] for e in count_map.items() if e[1] == max_count]

    def treeSum(self, root: TreeNode, count_map) -> int:
        sum = 0
        if root:
            left = self.treeSum(root.left, count_map)
            right = self.treeSum(root.right, count_map)
            sum = left + right + root.val
            count_map[sum] = count_map.get(sum, 0) + 1
        return sum
