import collections
from typing import List


class ListNode:
    pass


class Solution:

    # 82. 删除排序链表中的重复元素 II
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        # 思路：一次判断节点是否与前一个节点值相同，如果相同，则不加进来，先找到第一个节点。
        res, pre_val = ListNode(-101), 101
        pre = res
        while head:
            # 什么时候加入结果呢，与前一个值不相等，并且与后一个值也不相等的情况下加入
            if head.val != pre_val and (head.next is None or head.next.val != head.val):
                pre.next = head
                pre = pre.next
            pre_val = head.val
            head = head.next
            # 断开原来的联系
            pre.next = None
        return res.next

    # 594. 最长和谐子序列
    def findLHS(self, nums: List[int]) -> int:
        counter = collections.Counter()
        for num in nums:
            counter[num] += 1
        res = 0
        for k in counter.keys():
            if k - 1 in counter.keys():
                res = max(res, counter[k - 1] + counter[k])
            if k + 1 in counter.keys():
                res = max(res, counter[k + 1] + counter[k])
        return res

    # 946. 验证栈序列
    def validateStackSequences(self, pushed: List[int], popped: List[int]) -> bool:
        # 思路：先看popped的第一个元素i，然后pushed元素一次入栈，直到遇到i，弹出i。
        # 开始下一个元素i,如果i在栈顶或后面，则弹出或者一次压入直到该元素，否则返回false
        stack = []
        stack_set = set()
        pushed_index = 0
        for i in popped:
            if len(stack) > 0 and stack[len(stack) - 1] == i:
                stack.pop()
            elif i in stack_set:
                return False
            else:
                while pushed_index < len(pushed) and pushed[pushed_index] != i:
                    stack.append(pushed[pushed_index])
                    stack_set.add(pushed[pushed_index])
                    pushed_index += 1
                pushed_index += 1
        return True

    # 921. 使括号有效的最少添加
    def minAddToMakeValid(self, s: str) -> int:
        # 思路：从左往右，依次添加括号，对于已经完成匹配的，删除掉。如果前面是)，则添加对应的（
        # 如果对应的是（，则继续往前走，并记录)的数量。如果一旦达到平衡，则去掉，如果达不到平衡，则在最后的时候补充上去
        res = 0
        left = 0
        for w in s:
            if w == '(':
                left += 1
            elif left > 0:
                left -= 1
            else:
                res += 1
        return res + left

    # 905. 按奇偶排序数组
    def sortArrayByParity(self, nums: List[int]) -> List[int]:
        # list(filter(lambda num : num % 2 == 0,nums)) + list(filter(lambda  num : num % 2,nums))
        even, odd = 0, len(nums) - 1
        while even < odd:
            while even < len(nums) and nums[even] % 2 == 0:
                even += 1
            while odd >= 0 and nums[odd] % 2 == 1:
                odd -= 1
            if even < odd:
                nums[even], nums[odd] = nums[odd], nums[even]
        return nums
