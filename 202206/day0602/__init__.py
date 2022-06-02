from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    # 450. 删除二叉搜索树中的节点
    def deleteNode(self, root: Optional[TreeNode], key: int, parent=None, is_left=True) -> Optional[TreeNode]:
        # 思路：先寻找这个key,找到后，把它的左边 单独拿出来，把自己删了，把的右子树放在parent上，将左边追加到 右子树上
        # 如果不存在右子树呢？parent.left = root.left即可。
        if root:
            if root.val > key:
                self.deleteNode(root.left, key, root, True)
            elif root.val < key:
                self.deleteNode(root.right, key, root, False)
            else:
                # 找到相等的了，
                if root.right:
                    left = root.left
                    min_node = self.minNode(root.right)
                    min_node.left = left
                    if parent:
                        if is_left:
                            parent.left = root.right
                        else:
                            parent.right = root.right
                    else:
                        return root.right
                else:
                    if parent:
                        if is_left:
                            parent.left = root.left
                        else:
                            parent.right = root.left
                    else:
                        return root.left
        return root

    def minNode(self, root):
        if not root.left:
            return root
        return self.minNode(root.left)
