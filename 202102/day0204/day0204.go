package main

func main() {

}

//每日一题：643. 子数组最大平均数 I
func findMaxAverage(nums []int, k int) float64 {

	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	res := sum

	for i := k; i < len(nums); i++ {

		sum = sum - nums[i-k] + nums[i]
		if sum > res {
			res = sum
		}

	}

	return float64(res) / float64(k)
}
