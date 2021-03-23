package main

func main() {

}

//面试题 08.01. 三步问题
func waysToStep(n int) int {
	if n <= 2 {
		return n
	}
	// return waysToStep(n - 3)  + waysToStep(n - 2) + waysToStep(n - 1)

	//采用数组
	dp := make([]int, n+1)
	dp[1], dp[2], dp[3] = 1, 2, 4
	for i := 4; i < n+1; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + dp[i-3]) % 1000000007
	}
	return dp[n]
}
