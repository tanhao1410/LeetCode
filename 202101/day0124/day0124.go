package main

func main() {

}

//每日一题：674. 最长连续递增序列
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 1
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}
