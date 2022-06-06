from typing import List


class Solution:
    # 2094. 找出 3 位偶数
    def findEvenNumbers(self, digits: List[int]) -> List[int]:
        count = [0 for _ in range(10)]
        for i in digits:
            count[i] += 1

        def can_create_num(num):
            a, b, c = num % 10, num // 10 % 10, num // 100 % 10
            count[a] -= 1
            count[b] -= 1
            count[c] -= 1
            res = count[a] >= 0 and count[b] >= 0 and count[c] >= 0
            count[a] += 1
            count[b] += 1
            count[c] += 1
            return res

        return list(filter(can_create_num, range(100, 999, 2)))

    # LCP 01. 猜数字
    def game(self, guess: List[int], answer: List[int]) -> int:
        res = 0
        for i, j in zip(guess, answer):
            if i == j:
                res += 1
        return res


class MyCalendarThree:
    '''732. 我的日程安排表 III'''

    def __init__(self):
        self.items = {}

    def book(self, start: int, end: int) -> int:
        self.items[start] = self.items.get(start, 0) + 1
        self.items[end] = self.items.get(end, 0) - 1
        res = count = 0
        for _, i in sorted(self.items.items(), key=lambda e: e[0]):
            count += i
            res = max(res, count)
        return res
