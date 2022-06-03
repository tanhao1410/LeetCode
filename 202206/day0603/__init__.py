from typing import List


class Solution:
    # 885. 螺旋矩阵 III
    def spiralMatrixIII(self, rows: int, cols: int, rStart: int, cStart: int) -> List[List[int]]:
        res = [[rStart, cStart]]

        def append(r, c):
            if 0 <= c < cols and 0 <= r < rows:
                res.append([r, c])

        step = 1
        while len(res) < rows * cols:

            # 向右走
            for _ in range(step):
                cStart += 1
                append(rStart, cStart)
            # 向下走
            for _ in range(step):
                rStart += 1
                append(rStart, cStart)
            # 向左走
            step += 1
            for _ in range(step):
                cStart -= 1
                append(rStart, cStart)
            # 向上走
            for _ in range(step):
                rStart -= 1
                append(rStart, cStart)
            step += 1
        return res

    # 829. 连续整数求和
    def consecutiveNumbersSum(self, n: int) -> int:
        # 思路：一个，两个，三个，四个数，五个数。。。
        # 两个数:N = 2 * n + 1 and n > 0
        # 三个数:N = 3 * n and n > 1;
        # 四个数:N = 4n + 6 and n > 1;
        # ...... N = x* n + 是否偶数（x * (x - 1) / 2）
        res = 1
        count = 2
        while count * (count - 1) // 2 < n:
            if count % 2 == 0:
                if (n - (count - 1) * count // 2) % count == 0 and (n - (count - 1) * count // 2) // count > 0:
                    res += 1
            else:
                if n % count == 0 and n // count > 0:
                    res += 1
            count += 1
        return res

    # 690. 员工的重要性
    def getImportance(self, employees: List['Employee'], id: int) -> int:
        id_employ = {}
        for employ in employees:
            id_employ[employ.id] = employ
        return self.getImprtance2(id_employ, id)

    def getImprtance2(self, employees, id) -> int:
        res = employees[id].importance
        subordinates = employees[id].subordinates
        for sub in subordinates:
            res += self.getImprtance2(employees, id)
        return res


class Employee:
    def __init__(self, id: int, importance: int, subordinates: List[int]):
        self.id = id
        self.importance = importance
        self.subordinates = subordinates
