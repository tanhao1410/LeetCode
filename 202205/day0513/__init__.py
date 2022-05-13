class Solution:
    # 991. 坏了的计算器
    def brokenCalc(self, startValue: int, target: int) -> int:
        # 广度遍历+剪枝
        # 如果target>start,则所有小于start的都去除掉。如果target比目标小呢，直接返回即可。
        # 如果当前值的2倍小于目标值，则只需要加入2倍的，而不用加入-1的
        # if target <= startValue:
        #     return startValue - target
        # reached = {startValue}
        # q = deque(reached)
        # res = 0
        # while True:
        #     q_len = len(q)
        #     for _ in range(q_len):
        #         num = q.popleft()
        #         if num == target:
        #             return res
        #         if num * 2 not in reached:
        #             q.append(num * 2)
        #             reached.add(num * 2)
        #         if num * 2 > target and num - 1 > startValue and num - 1 not in reached:
        #             q.append(num - 1)
        #             reached.add(num - 1)
        #     res += 1
        res = 0
        while target > startValue:
            res += 1
            if target % 2 == 1:
                target += 1
            else:
                target //= 2
        return res + startValue - target

    # 984. 不含 AAA 或 BBB 的字符串
    def strWithout3a3b(self, a: int, b: int) -> str:
        res = ''
        # 思路：每一次追加一个字母，追加谁呢，剩余的ab谁多加谁，而且不能出现连着三个。
        while a > 0 or b > 0:
            if a > b:
                if len(res) > 1 and res[-1] == res[-2] and res[-1] == 'a':
                    res += 'b'
                    b -= 1
                else:
                    res += 'a'
                    a -= 1
            else:
                if len(res) > 1 and res[-1] == res[-2] and res[-1] == 'b':
                    res += 'a'
                    a -= 1
                else:
                    res += 'b'
                    b -= 1
        return res

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
