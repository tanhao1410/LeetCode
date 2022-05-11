class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


from collections import deque;


# 449. 序列化和反序列化二叉搜索树
class Codec:

    def serialize(self, root: TreeNode) -> str:
        """Encodes a tree to a single string.
        """
        res = ''
        # 如果有子节点，那么子节点入队，如果无子节点，则空入队。如果本身是空，添加一个nil,跳过
        q = deque()
        q.append(root)
        while len(q) > 0:
            head = q.popleft()
            if not head:
                res += 'nil'
            else:
                res += str(head.val)
                q.append(head.left)
                q.append(head.right)
            res += ','
        return res[:-1]

    def deserialize(self, data: str) -> TreeNode:
        """Decodes your encoded data to tree.
        """
        if data == 'nil':
            return None
        node_strs = data.split(',')
        index = 0
        head = TreeNode(int(node_strs[index]))
        index += 1
        q = deque()
        q.append(head)
        # 队列中有几个元素，就取出来对应2倍的字符串元素，放入其左右子树，如果左右子树不为空，加入队列
        while len(q) > 0:
            q_len = len(q)
            for _ in range(q_len):
                first = q.popleft()
                left = node_strs[index]
                if left != 'nil':
                    first.left = TreeNode(int(left))
                    q.append(first.left)
                index += 1
                right = node_strs[index]
                if right != 'nil':
                    first.right = TreeNode(int(right))
                    q.append(first.right)
                index += 1
        return head
