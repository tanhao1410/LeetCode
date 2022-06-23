class Solution:
    # 1387. 将整数按权重排序
    def getKth(self, lo: int, hi: int, k: int) -> int:
        map_count = {1 << i: i for i in range(32)}
        return sorted(range(lo, hi + 1), key=lambda i: self.get_count(i, map_count))[k - 1]

    def get_count(self, num, map_count) -> int:
        if num in map_count:
            return map_count[num]
        if num % 2 == 0:
            res = self.get_count(num // 2, map_count)
        else:
            res = self.get_count(num * 3 + 1, map_count)
        map_count[num] = res + 1
        return map_count[num]
