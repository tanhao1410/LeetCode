package main

func main() {

}

//746. 使用最小花费爬楼梯
func minCostClimbingStairs(cost []int) int {
	//思路：采用动态规划的方式dp[i]为从i处走到顶点所花费的，
	dp := make([]int, len(cost))
	dp[len(cost)-1] = cost[len(cost)-1] // 倒数第二个的最小花费是确定的
	dp[len(cost)-2] = cost[len(cost)-2] // 倒数第二个的最小花费是确定的
	for i := len(cost) - 3; i >= 0; i-- {
		if dp[i+1] > dp[i+2] {
			dp[i] = cost[i] + dp[i+2]
		} else {
			dp[i] = cost[i] + dp[i+1]
		}
	}
	if dp[0] > dp[1] {
		return dp[1]
	}
	return dp[0]
}
