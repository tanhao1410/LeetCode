package main

func main() {

}

//每日一题：724. 寻找数组的中心索引
func pivotIndex(nums []int) int {
	//先求总和，然后从左边，依次求和，若和+下一个数==总和-所有，则找到了
	sum := 0
	for _, v := range nums {
		sum += v
	}

	preSum := 0
	for i := 0; i < len(nums); i++ {
		//前面的数的和
		if i > 0 {
			preSum += nums[i-1]
		}
		//后面数的和
		tailSum := sum - preSum - nums[i]
		if tailSum == preSum {
			return i
		}
	}
	return -1
}
