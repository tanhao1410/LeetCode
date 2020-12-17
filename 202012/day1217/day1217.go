package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9, 3, 2, 8, 4, 9}, 4))
}

//每日一题：714. 买卖股票的最佳时机含手续费
func maxProfit(prices []int, fee int) int {
	//思路：采用动态规划表，dp[i]代表price[i:]能挣的最多的钱数
	//l[i]代表，dp[i]购买的起点价
	dp, l := make([]int, len(prices)), make([]int, len(prices))
	max := prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		//即有钱赚了
		if l[i+1] != 0 {
			if prices[i] < l[i+1] || prices[i]+fee < max {
				//比它后面购买时还要小，或比最大值要小
				//还要看max
				if max-prices[i]-fee > l[i+1]-prices[i] {
					dp[i] = dp[i+1] + max - prices[i] - fee
				} else {
					dp[i] = dp[i+1] - prices[i] + l[i+1]
				}
				l[i] = prices[i]
				max = 0
			} else {

				if prices[i] > max {
					max = prices[i]
				}
				dp[i] = dp[i+1]
				l[i] = l[i+1]
			}
		} else {
			//还没有赚钱时
			if prices[i] >= max {
				max = prices[i]
			} else {
				//比最大值要小
				if max-prices[i] > fee {
					dp[i] = max - prices[i] - fee
					l[i] = prices[i]
					max = 0
				}
			}
		}
	}
	return dp[0]
}
