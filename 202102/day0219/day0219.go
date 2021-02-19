package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(thirdMax([]int{1, 2, -2147483648}))
}

//485. 最大连续 1 的个数
func findMaxConsecutiveOnes(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == 1 {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 0
		}
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}

//414. 第三大的数
func thirdMax(nums []int) int {

	first, second, third := nums[0], math.MinInt32, math.MinInt32

	secondFlag, thirdFlag := false, false
	for i := 1; i < len(nums); i++ {
		if nums[i] == first || (nums[i] == second && secondFlag) || (nums[i] == third && thirdFlag) {
			continue
		}
		if nums[i] > first {
			first, second, third = nums[i], first, second
			if secondFlag {
				thirdFlag = true
			}
			secondFlag = true
		} else if nums[i] > second {
			second, third = nums[i], second
			if secondFlag {
				thirdFlag = true
			}
			secondFlag = true
		} else if nums[i] > third || (!thirdFlag && nums[i] == third) {
			third = nums[i]
			thirdFlag = true
		}

	}
	if thirdFlag {
		return third
	}
	return first
}
