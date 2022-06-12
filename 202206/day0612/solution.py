from typing import List, Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
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
