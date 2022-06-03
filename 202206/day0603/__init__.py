class Solution:
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
