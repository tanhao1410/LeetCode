class Solution:
    # 面试题 01.05. 一次编辑
    def oneEditAway(self, first: str, second: str) -> bool:
        def moreThanOne(first: str, second: str) -> bool:
            i, j, dif = 0, 0, 0
            while i < len(first) and j < len(second):
                if first[i] != second[j]:
                    i += 1
                    dif += 1
                else:
                    i += 1
                    j += 1
            return dif <= 1

        dif = 0
        if len(first) == len(second):
            for i in range(len(first)):
                if first[i] != second[i]:
                    dif += 1
            return dif <= 1
        elif len(first) - 1 == len(second):
            return moreThanOne(first, second)
        elif len(second) - 1 == len(first):
            return moreThanOne(second, first)
        return False
