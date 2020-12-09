package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(1, 1))
}

//343. 整数拆分
func integerBreak(n int) int {
	// n 不小于 2 且不大于 58。
	//特殊处理n = 2,n = 3 情况
	if n < 4 {
		return n - 1
	}
	dp := []int{1, 1, 2, 3}
	for i := 4; i < n+1; i++ {
		max := 0
		for j := 2; j < i; j++ {
			if dp[j]*dp[i-j] > max {
				max = dp[j] * dp[i-j]
			}
		}
		dp = append(dp, max)
	}
	return dp[n]
}

//每日一题：62. 不同路径
func uniquePaths(m int, n int) int {
	//思路：动态规划做法
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			//如果它下面可以走的话
			if i < m-1 {
				dp[i][j] += dp[i+1][j]
			}
			if j < n-1 {
				dp[i][j] += dp[i][j+1]
			}
		}
	}
	return dp[0][0]
}
