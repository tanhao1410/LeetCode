from typing import List, Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    # 1145. 二叉树着色游戏
    def btreeGameWinningMove(self, root: Optional[TreeNode], n: int, x: int) -> bool:
        # 1.选择x的父节点
        if root.val != x:
            # 求某一个节点下的所有节点数
            x_count = self.btree_node_count(root, x)
            if x_count * 2 < n:
                return True
        # 2.求子节点
        child_max = self.btree_child_node_max(root, x)
        if child_max * 2 > n:
            return True
        return False

    def btree_child_node_max(self, root: Optional[TreeNode], x: int):
        res = 0
        if root:
            if root.val == x:
                left = self.btree_node_count(root.left, contains=True)
                right = self.btree_node_count(root.right, contains=True)
                res = max(left, right)
            else:
                left = self.btree_child_node_max(root.left, x)
                right = self.btree_child_node_max(root.right, x)
                res = max(left, right)
        return res

    def btree_node_count(self, root: Optional[TreeNode], x=0, contains=False):
        res = 0
        if root:
            if contains or root.val == x:
                res += 1
                res += self.btree_node_count(root.left, contains=True)
                res += self.btree_node_count(root.right, contains=True)
            else:
                left = self.btree_node_count(root.left, x)
                right = self.btree_node_count(root.right, x)
                res = max(left, right)
        return res

    # 1008. 前序遍历构造二叉搜索树
    def bstFromPreorder(self, preorder: List[int]) -> Optional[TreeNode]:
        # 对preorder进行切分
        if preorder:
            root = TreeNode(val=preorder[0])
            split = len(preorder)
            for i in range(1, len(preorder)):
                if preorder[i] > preorder[0]:
                    split = i
                    break
            root.left = self.bstFromPreorder(preorder[1:split])
            root.right = self.bstFromPreorder(preorder[split:])
            return root

    # 890. 查找和替换模式
    def findAndReplacePattern(self, words: List[str], pattern: str) -> List[str]:
        res = []
        for word in words:
            # 判断是否符合
            match_dict = dict(zip(word, pattern))
            match_dict2 = dict(zip(pattern, word))
            if all(map(lambda e: match_dict[e[0]] == e[1] and match_dict2[e[1]] == e[0], zip(word, pattern))):
                res.append(word)
        return res
