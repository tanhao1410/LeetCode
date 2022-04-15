from typing import List, Optional


class NestedInteger:
    pass


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:

    # 21. 合并两个有序链表
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        head = ListNode()
        p = head
        while list2 is not None and list1 is not None:
            if list1.val > list2.val:
                p.next = list2
                list2 = list2.next
            else:
                p.next = list1
                list1 = list1.next
            p = p.next
        if list1 is None:
            p.next = list2
        else:
            p.next = list1
        return head.next

    # 385.迷你语法分析器
    def deserialize(self, s: str) -> NestedInteger:
        # 思路：如果是单独的数字，直接生成一个返回即可。
        # 如果不是数字，说明是list，将list里的每一个元素按照递归形式生成对象，放入里面即可。
        res = NestedInteger()
        if s == '[]':
            return res
        elif s.startswith('['):
            list = self.process(s)
            for ele in list:
                res.add(self.deserialize(ele))
        else:
            res.setInteger(int(s))
        return res

    # 将[1,2,[2,3]]拆分成list，
    def process(self, s: str) -> List:
        # 先把左右两边的括号去除掉
        res = []
        # 思路：用一个变量表示前面有多少【,如果前面【是零个，一旦读到,或读到结尾则生成一个新的
        # 否则，继续往下读。
        flag = 0
        start = 1
        end = 1
        while end < len(s) - 1:
            # 如果开始是[，说明内部元素是列表
            if s[end] == '[':
                flag += 1
            elif s[end] == ']':
                flag -= 1
            elif s[end] == ',' and flag == 0:
                # 此时需要生成一个元素
                res.append(s[start:end])
                start = end + 1
            end += 1
        # 最后一个元素加入进来
        res.append(s[start:end])
        return res
