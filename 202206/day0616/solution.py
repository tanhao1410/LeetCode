from typing import List
from collections import Counter


class Solution:
    # 1848. 到目标元素的最小距离
    def getMinDistance(self, nums: List[int], target: int, start: int) -> int:
        # 以start为中心向两边扩散，找到了target就返回
        for i in range(len(nums)):
            if (start + i < len(nums) and nums[start + i] == target) or (start - i >= 0 and nums[start - i] == target):
                return i

    # 1169. 查询无效交易
    def invalidTransactions(self, transactions: List[str]) -> List[str]:
        # 先按交易名稱進行分类，名称->交易(城市，时间，金额)
        name_map = {}
        for transaction in transactions:
            items = transaction.split(',')
            name_list = name_map.get(items[0], [])
            item = [transaction]
            item.extend(items[1:])
            name_list.append(item)
            name_map[items[0]] = name_list
        res = []
        for name_list in name_map.values():
            for i in range(len(name_list)):
                if int(name_list[i][2]) > 1000:
                    res.append(name_list[i][0])
                    continue
                for j in range(len(name_list)):
                    if i != j and abs(int(name_list[i][1]) - int(name_list[j][1])) <= 60 and name_list[i][3] != \
                            name_list[j][3]:
                        res.append(name_list[i][0])
                        break
        return res

    # 532. 数组中的 k-diff 数对
    def findPairs(self, nums: List[int], k: int) -> int:
        nums_counter = Counter(nums)
        if k == 0:
            return sum(map(lambda e: 1, filter(lambda e: e > 1, nums_counter.values())))
        nums = nums_counter.keys()
        return sum([1 for v in nums if v + k in nums])
