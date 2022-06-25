from typing import List


class Solution:

    def maxCoins(self, piles: List[int]) -> int:
        piles = sorted(piles)[len(piles) // 3:]
        return sum([piles[n] for n in range(len(piles)) if n % 2 == 0])

    # 剑指 Offer II 091. 粉刷房子
    def minCost(self, costs: List[List[int]]) -> int:
        dp = [[n for n in x] for x in costs]
        for i in range(1, len(dp)):
            for y in range(3):
                min_cost = 1000000
                for pre_y in range(3):
                    if y != pre_y:
                        min_cost = min(min_cost, dp[i - 1][pre_y])
                dp[i][y] += min_cost
        return min(dp[len(dp) - 1])
