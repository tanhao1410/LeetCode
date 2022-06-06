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
