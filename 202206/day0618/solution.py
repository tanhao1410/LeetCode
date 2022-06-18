from typing import List


class Node:
    def __init__(self, val=None, next=None):
        self.val = val
        self.next = next


class Solution:
    # 1828. 统计一个圆中点的数目
    def countPoints(self, points: List[List[int]], queries: List[List[int]]) -> List[int]:
        def distace_double(p1, p2):
            return (p1[0] - p2[0]) ** 2 + (p1[1] - p2[1]) ** 2

        res = []
        for query in queries:
            count = 0
            for point in points:
                if distace_double(query, point) <= query[2] * query[2]:
                    count += 1
            res.append(count)
        return res

    # 剑指 Offer II 029. 排序的循环链表
    def insert(self, head: 'Node', insertVal: int) -> 'Node':
        new_node = Node(val=insertVal)
        new_node.next = new_node
        # 存在一种可能性，即所有值都相等。否则，在开始处，必然存在一个降序
        # 列表为空，直接返回即可。列表就一个的情况，即所有值都相等
        if head:
            # 判断是否所有的都相等。循环找到下降的，或者走到了原点
            p = head.next
            while p != head:
                # 找到插入位置
                if insertVal >= p.val and insertVal <= p.next.val:
                    p_next = p.next
                    p.next = new_node
                    new_node.next = p_next
                    return head
                if p.next.val < p.val:
                    # 如果p的下一个结点比自己还要小，说明到达了起点。
                    if insertVal <= p.next.val or insertVal >= p.val:
                        # 可以插入到头部
                        first = p.next
                        p.next = new_node
                        new_node.next = first
                        return head
                p = p.next
            # 走到这，说明所有节点都相等，随便找一个位置插入即可
            head_next = head.next
            head.next = new_node
            new_node.next = head_next
            return head
        return new_node
