package main

func main() {

}

//每日一题：509. 斐波那契数
func fib(n int) int {
	dp := make([]int, n+1)
	if n == 0 {
		return 0
	}
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
